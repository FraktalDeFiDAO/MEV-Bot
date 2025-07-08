package main

import (
	"context"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/joho/godotenv"

	"github.com/FraktalDeFiDAO/MEV-Bot/pkg/ethutil"
	"github.com/FraktalDeFiDAO/MEV-Bot/pkg/watcher"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

)

type runner interface{ Run(context.Context) error }

// connectClient abstracts ethutil.ConnectClient for testability.
var (
	connectClient   = ethutil.ConnectClient
	newBlockWatcher = func(sub watcher.HeaderSubscriber) runner { return watcher.NewBlockWatcher(sub) }
	tradeABI        abi.ABI
	tradeEventID    common.Hash
	newEventWatcher = func(sub watcher.LogSubscriber, q ethereum.FilterQuery) runner {
		return watcher.NewEventWatcher(sub, q, profitLogHandler)
	}
)

func init() {
	var err error
	tradeABI, err = abi.JSON(strings.NewReader(`[{"anonymous":false,"inputs":[{"indexed":false,"internalType":"uint256","name":"amountIn","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"profit","type":"uint256"}],"name":"TradeExecuted","type":"event"}]`))
	if err != nil {
		panic(err)
	}
	tradeEventID = crypto.Keccak256Hash([]byte("TradeExecuted(uint256,uint256)"))
}

// Entry point for the MEV bot. Connects to an Arbitrum node and listens for events.

func run(ctx context.Context, rpcURL string) error {
	client, err := connectClient(ctx, rpcURL)
	if err != nil {
		return err
	}
	defer client.Close()

	if _, err := client.ChainID(ctx); err != nil {
		return err
	}

	log.Println("connected to arbitrum", rpcURL)

	bw := newBlockWatcher(client)
	ew := newEventWatcher(client, ethereum.FilterQuery{})

	// run watchers until context cancellation
	go func() {
		if err := bw.Run(ctx); err != nil && ctx.Err() == nil {
			log.Printf("block watcher error: %v", err)
		}
	}()
	go func() {
		if err := ew.Run(ctx); err != nil && ctx.Err() == nil {
			log.Printf("event watcher error: %v", err)
		}
	}()

	<-ctx.Done()
	return ctx.Err()
}

// profitLogHandler decodes TradeExecuted events and prints the profit.
func profitLogHandler(l types.Log) {
	if len(l.Topics) > 0 && l.Topics[0] == tradeEventID {
		var ev struct {
			AmountIn *big.Int
			Profit   *big.Int
		}
		if err := tradeABI.UnpackIntoInterface(&ev, "TradeExecuted", l.Data); err == nil {
			log.Printf("profit event input=%s profit=%s tx=%s", ev.AmountIn.String(), ev.Profit.String(), l.TxHash.Hex())
		} else {
			log.Printf("profit event decode error: %v", err)
		}
	} else {
		log.Printf("log tx: %s", l.TxHash.Hex())
	}
}

func main() {
	// load environment variables from .env if present
	_ = godotenv.Load()

	if os.Getenv("DEBUG") != "" {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("debug logging enabled")
	} else {
		log.SetFlags(log.LstdFlags)
	}

	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		rpcURL = "https://arb1.arbitrum.io/rpc"
	}

	if err := run(context.Background(), rpcURL); err != nil {
		log.Fatalf("failed to connect to arbitrum rpc: %v", err)
	}
}
