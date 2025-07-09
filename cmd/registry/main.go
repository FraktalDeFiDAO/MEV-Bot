package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/FraktalDeFiDAO/MEV-Bot/pkg/registry"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
	fmt.Println("  add-pool <pool> [t0 t1 [id]]  register a pool")
	os.Exit(1)
}

func connect(ctx context.Context) (*registry.Client, *ethclient.Client, error) {
	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		rpcURL = "https://arb1.arbitrum.io/rpc"
	}
	regAddr := os.Getenv("REGISTRY_ADDRESS")
	keyHex := os.Getenv("PRIVATE_KEY")
	if regAddr == "" || keyHex == "" {
		return nil, nil, fmt.Errorf("REGISTRY_ADDRESS and PRIVATE_KEY must be set")
	}
	rpc, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		return nil, nil, err
	}
	key, err := crypto.HexToECDSA(strings.TrimPrefix(keyHex, "0x"))
	if err != nil {
		return nil, nil, err
	}
	client, err := registry.New(ctx, common.HexToAddress(regAddr), rpc, key)
	if err != nil {
		return nil, nil, err
	}
	return client, rpc, nil}

func main() {
	if len(os.Args) < 2 {
		usage()
	}
	ctx := context.Background()
	client, rpc, err := connect(ctx)
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
		if len(os.Args) < 3 {
			usage()
		}
		pool := common.HexToAddress(os.Args[2])
		var t0, t1 common.Address
		var id uint64
		if len(os.Args) >= 5 {
			t0 = common.HexToAddress(os.Args[3])
			t1 = common.HexToAddress(os.Args[4])
			if len(os.Args) > 5 {
				if v, err := strconv.ParseUint(os.Args[5], 10, 64); err == nil {
					id = v
				}
			}
		} else {
			pairABI, _ := abi.JSON(strings.NewReader(`[{"inputs":[],"name":"token0","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"token1","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"}]`))
			bound := bind.NewBoundContract(pool, pairABI, rpc, rpc, rpc)
			var out0, out1 []interface{}
			if err := bound.Call(&bind.CallOpts{Context: ctx}, &out0, "token0"); err != nil {
				log.Fatal(err)
			}
			if err := bound.Call(&bind.CallOpts{Context: ctx}, &out1, "token1"); err != nil {
				log.Fatal(err)
			}
			t0 = out0[0].(common.Address)
			t1 = out1[0].(common.Address)
		}
		tx, err := client.AddToken(t0, 18)
		if err == nil {
			client.WaitMined(ctx, tx)
		}
		tx, err = client.AddToken(t1, 18)
		if err == nil {
			client.WaitMined(ctx, tx)
		}
		tx, err = client.AddPool(pool, t0, t1, id)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(tx.Hash().Hex())
	default:
		usage()
	}
}
