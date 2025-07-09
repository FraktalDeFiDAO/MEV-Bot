package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/FraktalDeFiDAO/MEV-Bot/pkg/registry"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func usage() {
	fmt.Println("Usage:")
	fmt.Println("  registry <command> [args]")
	fmt.Println("Commands:")
	fmt.Println("  tokens                   list registered tokens")
	fmt.Println("  pools                    list registered pools")
	fmt.Println("  add-token <addr> [dec]   register a token")
	fmt.Println("  add-pool <pool> <t0> <t1> [id]  register a pool")
	os.Exit(1)
}

func connect(ctx context.Context) (*registry.Client, error) {
	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		rpcURL = "https://arb1.arbitrum.io/rpc"
	}
	regAddr := os.Getenv("REGISTRY_ADDRESS")
	keyHex := os.Getenv("PRIVATE_KEY")
	if regAddr == "" || keyHex == "" {
		return nil, fmt.Errorf("REGISTRY_ADDRESS and PRIVATE_KEY must be set")
	}
	rpc, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		return nil, err
	}
	key, err := crypto.HexToECDSA(strings.TrimPrefix(keyHex, "0x"))
	if err != nil {
		return nil, err
	}
	return registry.New(ctx, common.HexToAddress(regAddr), rpc, key)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}
	ctx := context.Background()
	client, err := connect(ctx)
	if err != nil {
		log.Fatalf("connect: %v", err)
	}
	switch os.Args[1] {
	case "tokens":
		toks, err := client.Tokens(ctx)
		if err != nil {
			log.Fatal(err)
		}
		for _, t := range toks {
			fmt.Println(t.Hex())
		}
	case "pools":
		pools, err := client.Pools(ctx)
		if err != nil {
			log.Fatal(err)
		}
		for _, p := range pools {
			info, err := client.PoolInfo(ctx, p)
			if err != nil {
				log.Println("info", err)
				continue
			}
			fmt.Printf("%s %s %s %s\n", p.Hex(), info.Token0.Hex(), info.Token1.Hex(), info.ExchangeID.String())
		}
	case "add-token":
		if len(os.Args) < 3 {
			usage()
		}
		addr := common.HexToAddress(os.Args[2])
		dec := uint8(18)
		if len(os.Args) > 3 {
			if v, err := strconv.Atoi(os.Args[3]); err == nil {
				dec = uint8(v)
			}
		}
		tx, err := client.AddToken(addr, dec)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(tx.Hash().Hex())
	case "add-pool":
		if len(os.Args) < 5 {
			usage()
		}
		pool := common.HexToAddress(os.Args[2])
		t0 := common.HexToAddress(os.Args[3])
		t1 := common.HexToAddress(os.Args[4])
		id := uint64(0)
		if len(os.Args) > 5 {
			if v, err := strconv.ParseUint(os.Args[5], 10, 64); err == nil {
				id = v
			}
		}
		tx, err := client.AddPool(pool, t0, t1, id)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(tx.Hash().Hex())
	default:
		usage()
	}
}
