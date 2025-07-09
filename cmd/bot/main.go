package main

import (
	"context"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/joho/godotenv"

	"github.com/FraktalDeFiDAO/MEV-Bot/pkg/arb"
	"github.com/FraktalDeFiDAO/MEV-Bot/pkg/ethutil"
	"github.com/FraktalDeFiDAO/MEV-Bot/pkg/market"
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
	syncABI         abi.ABI
	syncEventID     common.Hash
	pairABI         abi.ABI
	pairEventID     common.Hash
	poolABI         abi.ABI
	poolEventID     common.Hash
	newEventWatcher = func(sub watcher.LogSubscriber, q ethereum.FilterQuery) runner {
		return watcher.NewEventWatcher(sub, q, profitLogHandler)
	}
	newPairWatcher = func(sub watcher.LogSubscriber, q ethereum.FilterQuery) runner {
		return watcher.NewEventWatcher(sub, q, pairLogHandler)
	}
	arbMon      *arb.Monitor
	marketStore = market.New()
)

func init() {
	var err error
	tradeABI, err = abi.JSON(strings.NewReader(`[{"anonymous":false,"inputs":[{"indexed":false,"internalType":"uint256","name":"amountIn","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"profit","type":"uint256"}],"name":"TradeExecuted","type":"event"}]`))
	if err != nil {
		panic(err)
	}
	tradeEventID = crypto.Keccak256Hash([]byte("TradeExecuted(uint256,uint256)"))

	syncABI, err = abi.JSON(strings.NewReader(`[{"anonymous":false,"inputs":[{"indexed":false,"internalType":"uint112","name":"reserve0","type":"uint112"},{"indexed":false,"internalType":"uint112","name":"reserve1","type":"uint112"}],"name":"Sync","type":"event"}]`))
	if err != nil {
		panic(err)
	}
	syncEventID = crypto.Keccak256Hash([]byte("Sync(uint112,uint112)"))

	pairABI, err = abi.JSON(strings.NewReader(`[{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"token0","type":"address"},{"indexed":true,"internalType":"address","name":"token1","type":"address"},{"indexed":false,"internalType":"address","name":"pair","type":"address"},{"indexed":false,"internalType":"uint256","name":"","type":"uint256"}],"name":"PairCreated","type":"event"}]`))
	if err != nil {
		panic(err)
	}
	pairEventID = crypto.Keccak256Hash([]byte("PairCreated(address,address,address,uint256)"))

	poolABI, err = abi.JSON(strings.NewReader(`[{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"token0","type":"address"},{"indexed":true,"internalType":"address","name":"token1","type":"address"},{"indexed":false,"internalType":"uint24","name":"fee","type":"uint24"},{"indexed":false,"internalType":"int24","name":"tickSpacing","type":"int24"},{"indexed":false,"internalType":"address","name":"pool","type":"address"}],"name":"PoolCreated","type":"event"}]`))
	if err != nil {
		panic(err)
	}
	poolEventID = crypto.Keccak256Hash([]byte("PoolCreated(address,address,uint24,int24,address)"))
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
	// listen for TradeExecuted and Sync events
	query := ethereum.FilterQuery{
		Topics: [][]common.Hash{{tradeEventID, syncEventID}},
	}
	ew := newEventWatcher(client, query)

	// watch factory PairCreated/PoolCreated events if factories provided
	var pw runner
	if fenv := os.Getenv("FACTORIES"); fenv != "" {
		var addrs []common.Address
		for _, a := range strings.Split(fenv, ",") {
			addr := common.HexToAddress(strings.TrimSpace(a))
			addrs = append(addrs, addr)
		}
		pquery := ethereum.FilterQuery{
			Addresses: addrs,
			Topics:    [][]common.Hash{{pairEventID, poolEventID}},
		}
		pw = newPairWatcher(client, pquery)
	}

	if arbMon == nil {
		pairEnv := os.Getenv("PAIRS")
		var pairs [][2]common.Address
		for _, seg := range strings.Split(pairEnv, ";") {
			parts := strings.Split(seg, ",")
			if len(parts) == 2 {
				p0 := common.HexToAddress(strings.TrimSpace(parts[0]))
				p1 := common.HexToAddress(strings.TrimSpace(parts[1]))
				pairs = append(pairs, [2]common.Address{p0, p1})
			}
		}
		arbMon = arb.NewMonitor(pairs, 500, 1)
	}

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
	if pw != nil {
		go func() {
			if err := pw.Run(ctx); err != nil && ctx.Err() == nil {
				log.Printf("pair watcher error: %v", err)
			}
		}()
	}
	<-ctx.Done()
	return ctx.Err()
}

// profitLogHandler decodes TradeExecuted events and prints the profit.
func profitLogHandler(l types.Log) {
	if len(l.Topics) > 0 {
		switch l.Topics[0] {
		case tradeEventID:
			var ev struct {
				AmountIn *big.Int
				Profit   *big.Int
			}
			if err := tradeABI.UnpackIntoInterface(&ev, "TradeExecuted", l.Data); err == nil {
				log.Printf("profit event input=%s profit=%s tx=%s", ev.AmountIn.String(), ev.Profit.String(), l.TxHash.Hex())
			} else {
				log.Printf("profit event decode error: %v", err)
			}
		case syncEventID:
			var ev struct {
				Reserve0 *big.Int
				Reserve1 *big.Int
			}
			if err := syncABI.UnpackIntoInterface(&ev, "Sync", l.Data); err == nil {
				price := new(big.Float).Quo(new(big.Float).SetInt(ev.Reserve1), new(big.Float).SetInt(ev.Reserve0))
				f, _ := price.Float64()
				log.Printf("price update pool=%s price=%f tx=%s", l.Address.Hex(), f, l.TxHash.Hex())
				if arbMon != nil {
					arbMon.Update(l.Address, ev.Reserve0, ev.Reserve1)
				}
			} else {
				log.Printf("sync event decode error: %v", err)
			}
		default:
			log.Printf("log tx: %s", l.TxHash.Hex())
		}
	} else {
		log.Printf("log tx: %s", l.TxHash.Hex())
	}
}

// pairLogHandler collects pool addresses from factory events.
func pairLogHandler(l types.Log) {
	if len(l.Topics) == 0 {
		return
	}
	switch l.Topics[0] {
	case pairEventID:
		var ev struct {
			Token0 common.Address
			Token1 common.Address
			Pair   common.Address
			Arg3   *big.Int
		}
		if err := pairABI.UnpackIntoInterface(&ev, "PairCreated", l.Data); err == nil {
			log.Printf("pair created %s", ev.Pair.Hex())
			if marketStore != nil {
				marketStore.Add(ev.Pair)
				marketStore.AddToken(ev.Token0)
				marketStore.AddToken(ev.Token1)
			}
		} else {
			log.Printf("pair decode error: %v", err)
		}
	case poolEventID:
		var ev struct {
			Token0      common.Address
			Token1      common.Address
			Fee         *big.Int
			TickSpacing *big.Int
			Pool        common.Address
		}
		if err := poolABI.UnpackIntoInterface(&ev, "PoolCreated", l.Data); err == nil {
			log.Printf("pool created %s", ev.Pool.Hex())
			if marketStore != nil {
				marketStore.Add(ev.Pool)
				marketStore.AddToken(ev.Token0)
				marketStore.AddToken(ev.Token1)
			}
		} else {
			log.Printf("pool decode error: %v", err)
		}
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
