// Command registry exposes a simple CLI for interacting with the on-chain
// registry contract. It supports listing tokens and pools as well as adding
// new tokens or pools.
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
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	dialRPC     = ethclient.DialContext
	newRegistry = registry.New
)

type registryClient interface {
	AddPool(common.Address, common.Address, common.Address, uint64) (*types.Transaction, error)
	AddToken(common.Address, uint8) (*types.Transaction, error)
	WaitMined(context.Context, *types.Transaction) (*types.Receipt, error)
	Tokens(context.Context) ([]common.Address, error)
	Pools(context.Context) ([]common.Address, error)
	PoolInfo(context.Context, common.Address) (registry.PoolInfo, error)
}

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
	rpc, err := dialRPC(ctx, rpcURL)
	if err != nil {
		return nil, nil, err
	}
	key, err := crypto.HexToECDSA(strings.TrimPrefix(keyHex, "0x"))
	if err != nil {
		return nil, nil, err
	}
	client, err := newRegistry(ctx, common.HexToAddress(regAddr), rpc, key)
	if err != nil {
		return nil, nil, err
	}
	return client, rpc, nil
}

func handle(ctx context.Context, client registryClient, rpc *ethclient.Client, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("missing command")
	}
	switch args[0] {
	case "tokens":
		toks, err := client.Tokens(ctx)
		if err != nil {
			return err
		}
		for _, t := range toks {
			fmt.Println(t.Hex())
		}
		return nil
	case "pools":
		pools, err := client.Pools(ctx)
		if err != nil {
			return err
		}
		for _, p := range pools {
			info, err := client.PoolInfo(ctx, p)
			if err != nil {
				log.Println("info", err)
				continue
			}
			fmt.Printf("%s %s %s %s\n", p.Hex(), info.Token0.Hex(), info.Token1.Hex(), info.ExchangeID.String())
		}
		return nil
	case "add-token":
		if len(args) < 2 {
			return fmt.Errorf("add-token requires address")
		}
		addr := common.HexToAddress(args[1])
		dec := uint8(18)
		if len(args) > 2 {
			if v, err := strconv.Atoi(args[2]); err == nil {
				dec = uint8(v)
			}
		}
		tx, err := client.AddToken(addr, dec)
		if err != nil {
			return err
		}
		fmt.Println(tx.Hash().Hex())
		return nil
	case "add-pool":
		if len(args) < 2 {
			return fmt.Errorf("add-pool requires pool address")
		}
		pool := common.HexToAddress(args[1])
		var t0, t1 common.Address
		var id uint64
		if len(args) >= 4 {
			t0 = common.HexToAddress(args[2])
			t1 = common.HexToAddress(args[3])
			if len(args) > 4 {
				if v, err := strconv.ParseUint(args[4], 10, 64); err == nil {
					id = v
				}
			}
		} else {
			pairABI, _ := abi.JSON(strings.NewReader(`[{"inputs":[],"name":"token0","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"token1","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"}]`))
			bound := bind.NewBoundContract(pool, pairABI, rpc, rpc, rpc)
			var out0, out1 []interface{}
			if err := bound.Call(&bind.CallOpts{Context: ctx}, &out0, "token0"); err != nil {
				return err
			}
			if err := bound.Call(&bind.CallOpts{Context: ctx}, &out1, "token1"); err != nil {
				return err
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
			return err
		}
		fmt.Println(tx.Hash().Hex())
		return nil
	default:
		return fmt.Errorf("unknown command %s", args[0])
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}
	ctx := context.Background()
	client, rpc, err := connect(ctx)
	if err != nil {
		log.Fatalf("connect: %v", err)
	}
	if err := handle(ctx, client, rpc, os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}
