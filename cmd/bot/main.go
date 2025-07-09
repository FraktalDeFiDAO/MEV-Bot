package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"

	"github.com/FraktalDeFiDAO/MEV-Bot/pkg/arb"
	"github.com/FraktalDeFiDAO/MEV-Bot/pkg/ethutil"
	"github.com/FraktalDeFiDAO/MEV-Bot/pkg/market"
	"github.com/FraktalDeFiDAO/MEV-Bot/pkg/registry"
	"github.com/FraktalDeFiDAO/MEV-Bot/pkg/watcher"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type runner interface{ Run(context.Context) error }

type registryClient interface {
	AddPool(common.Address, common.Address, common.Address, uint64) (*types.Transaction, error)
	AddToken(common.Address, uint8) (*types.Transaction, error)
	Tokens(context.Context) ([]common.Address, error)
	Pools(context.Context) ([]common.Address, error)
	PoolInfo(context.Context, common.Address) (registry.PoolInfo, error)
}

// connectClient abstracts ethutil.ConnectClient for testability.
var (
	connectClient = ethutil.ConnectClient
	newRegistry   = func(ctx context.Context, addr common.Address, rpc *ethclient.Client, key *ecdsa.PrivateKey) (registryClient, error) {
		return registry.New(ctx, addr, rpc, key)
	}
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
	marketStore *market.Persistent
	regClient   registryClient
	knownTokens map[common.Address]struct{}
	knownPools  map[common.Address]struct{}
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

func run(ctx context.Context, rpcURL, regAddr, keyHex string) error {
	client, err := connectClient(ctx, rpcURL)
	if err != nil {
		return err
	}
	defer client.Close()

	if _, err := client.ChainID(ctx); err != nil {
		return err
	}

	log.Println("connected to arbitrum", rpcURL)
	if regAddr != "" && keyHex != "" {
		key, err := crypto.HexToECDSA(strings.TrimPrefix(keyHex, "0x"))
		if err == nil {
			rc, err := newRegistry(ctx, common.HexToAddress(regAddr), client, key)
			if err == nil {
				regClient = rc
			} else {
				log.Printf("registry init error: %v", err)
			}
		} else {
			log.Printf("key error: %v", err)
		}
	}

	if regClient != nil {
		loadRegistry(ctx)
	}
	syncRegistry()

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

	startServer(":8080")

	<-ctx.Done()
	return ctx.Err()
}

// syncRegistry registers any tokens loaded from disk with the on-chain registry.
func syncRegistry() {
	if regClient == nil || marketStore == nil {
		return
	}

	if knownTokens == nil {
		knownTokens = make(map[common.Address]struct{})
	}
	if knownPools == nil {
		knownPools = make(map[common.Address]struct{})
	}

  ctx := context.Background()

	onchainTokens, err := regClient.Tokens(ctx)
	if err != nil {
		log.Printf("registry token query error: %v", err)
	}
	for _, t := range onchainTokens {
		knownTokens[t] = struct{}{}
	}

	for _, t := range marketStore.ListTokens() {
		if _, ok := knownTokens[t]; ok {
			continue
		}
		if tx, err := regClient.AddToken(t, 18); err != nil {
			log.Printf("registry sync token error: %v", err)
		} else {
			log.Printf("registry add token %s tx=%s", t.Hex(), tx.Hash().Hex())
			knownTokens[t] = struct{}{}

		}
	}

	onchainPools, err := regClient.Pools(ctx)
	if err != nil {
		log.Printf("registry pool query error: %v", err)
	}
	for _, p := range onchainPools {
		knownPools[p] = struct{}{}
	}

	for _, p := range marketStore.ListPools() {
		if p.Token0 == (common.Address{}) || p.Token1 == (common.Address{}) {
			continue
		}
		if _, ok := knownPools[p.Address]; ok {
			continue
		}
		if tx, err := regClient.AddPool(p.Address, p.Token0, p.Token1, 0); err != nil {
			log.Printf("registry sync pool error: %v", err)
		} else {
			log.Printf("registry add pool %s tx=%s", p.Address.Hex(), tx.Hash().Hex())
			knownPools[p.Address] = struct{}{}
		}
	}
}

// loadRegistry fetches tokens and pools from the on-chain registry.
func loadRegistry(ctx context.Context) {
	if regClient == nil || marketStore == nil {
		return
	}
	toks, err := regClient.Tokens(ctx)
	if err == nil {
		for _, t := range toks {
			marketStore.AddToken(t)
			knownTokens[t] = struct{}{}
		}
	} else {
		log.Printf("registry token load error: %v", err)
	}
	pools, err := regClient.Pools(ctx)
	if err == nil {
		for _, p := range pools {
			info, err := regClient.PoolInfo(ctx, p)
			if err == nil {
				marketStore.AddPool(p, info.Token0, info.Token1)
				knownPools[p] = struct{}{}
			}
		}
	} else {
		log.Printf("registry pool load error: %v", err)
	}
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
				marketStore.AddPool(ev.Pair, ev.Token0, ev.Token1)
				marketStore.AddToken(ev.Token0)
				marketStore.AddToken(ev.Token1)
			}
			if regClient != nil {
				if _, ok := knownTokens[ev.Token0]; !ok {
					if tx, err := regClient.AddToken(ev.Token0, 18); err != nil {
						log.Printf("registry token0 error: %v", err)
					} else {
						log.Printf("registry add token %s tx=%s", ev.Token0.Hex(), tx.Hash().Hex())
						knownTokens[ev.Token0] = struct{}{}
					}
				}
				if _, ok := knownTokens[ev.Token1]; !ok {
					if tx, err := regClient.AddToken(ev.Token1, 18); err != nil {
						log.Printf("registry token1 error: %v", err)
					} else {
						log.Printf("registry add token %s tx=%s", ev.Token1.Hex(), tx.Hash().Hex())
						knownTokens[ev.Token1] = struct{}{}
					}
				}
				if _, ok := knownPools[ev.Pair]; !ok {
					if tx, err := regClient.AddPool(ev.Pair, ev.Token0, ev.Token1, 0); err != nil {
						log.Printf("registry pool error: %v", err)
					} else {
						log.Printf("registry add pool %s tx=%s", ev.Pair.Hex(), tx.Hash().Hex())
						knownPools[ev.Pair] = struct{}{}
					}

				}
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
				marketStore.AddPool(ev.Pool, ev.Token0, ev.Token1)
				marketStore.AddToken(ev.Token0)
				marketStore.AddToken(ev.Token1)
			}
			if regClient != nil {
				if _, ok := knownTokens[ev.Token0]; !ok {
					if tx, err := regClient.AddToken(ev.Token0, 18); err != nil {
						log.Printf("registry token0 error: %v", err)
					} else {
						log.Printf("registry add token %s tx=%s", ev.Token0.Hex(), tx.Hash().Hex())
						knownTokens[ev.Token0] = struct{}{}
					}
				}
				if _, ok := knownTokens[ev.Token1]; !ok {
					if tx, err := regClient.AddToken(ev.Token1, 18); err != nil {
						log.Printf("registry token1 error: %v", err)
					} else {
						log.Printf("registry add token %s tx=%s", ev.Token1.Hex(), tx.Hash().Hex())
						knownTokens[ev.Token1] = struct{}{}
					}
				}
				if _, ok := knownPools[ev.Pool]; !ok {
					if tx, err := regClient.AddPool(ev.Pool, ev.Token0, ev.Token1, 0); err != nil {
						log.Printf("registry pool error: %v", err)
					} else {
						log.Printf("registry add pool %s tx=%s", ev.Pool.Hex(), tx.Hash().Hex())
						knownPools[ev.Pool] = struct{}{}
					}

				}
			}
		} else {
			log.Printf("pool decode error: %v", err)
		}
	}
}

func startServer(addr string) {
	http.HandleFunc("/tokens", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if marketStore == nil {
			http.Error(w, "no store", http.StatusInternalServerError)
			return
		}
		switch r.Method {
		case http.MethodGet:
			json.NewEncoder(w).Encode(marketStore.ListTokens())
		case http.MethodPost:
			var in struct {
				Address string `json:"address"`
			}
			if json.NewDecoder(r.Body).Decode(&in) == nil {
				addr := common.HexToAddress(in.Address)
				marketStore.AddToken(addr)
				if regClient != nil {
					if _, ok := knownTokens[addr]; !ok {
						if tx, err := regClient.AddToken(addr, 18); err == nil {
							log.Printf("registry add token %s tx=%s", addr.Hex(), tx.Hash().Hex())
							knownTokens[addr] = struct{}{}
						} else {
							log.Printf("registry token error: %v", err)
						}
					}

				}
			}
			w.WriteHeader(http.StatusOK)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/pools", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if marketStore == nil {
			http.Error(w, "no store", http.StatusInternalServerError)
			return
		}
		switch r.Method {
		case http.MethodGet:
			json.NewEncoder(w).Encode(marketStore.ListPools())
		case http.MethodPost:
			var in struct {
				Address string `json:"address"`
				Token0  string `json:"token0"`
				Token1  string `json:"token1"`
			}
			if json.NewDecoder(r.Body).Decode(&in) == nil {
				p := common.HexToAddress(in.Address)
				t0 := common.HexToAddress(in.Token0)
				t1 := common.HexToAddress(in.Token1)
				marketStore.AddPool(p, t0, t1)
				if regClient != nil {
					if _, ok := knownTokens[t0]; !ok {
						if tx, err := regClient.AddToken(t0, 18); err == nil {
							log.Printf("registry add token %s tx=%s", t0.Hex(), tx.Hash().Hex())
							knownTokens[t0] = struct{}{}
						} else {
							log.Printf("registry token0 error: %v", err)
						}
					}
					if _, ok := knownTokens[t1]; !ok {
						if tx, err := regClient.AddToken(t1, 18); err == nil {
							log.Printf("registry add token %s tx=%s", t1.Hex(), tx.Hash().Hex())
							knownTokens[t1] = struct{}{}
						} else {
							log.Printf("registry token1 error: %v", err)
						}
					}
					if _, ok := knownPools[p]; !ok {
						if tx, err := regClient.AddPool(p, t0, t1, 0); err == nil {
							log.Printf("registry add pool %s tx=%s", p.Hex(), tx.Hash().Hex())
							knownPools[p] = struct{}{}
						} else {
							log.Printf("registry pool error: %v", err)
						}
					}

				}
			}
			w.WriteHeader(http.StatusOK)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	go func() {
		if err := http.ListenAndServe(addr, nil); err != nil {
			log.Printf("api server error: %v", err)
		}
	}()
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

	cachePath := os.Getenv("MARKET_CACHE")
	if cachePath == "" {
		cachePath = "market.json"
	}
	marketStore = market.LoadFromFile(cachePath)
	knownTokens = make(map[common.Address]struct{})
	knownPools = make(map[common.Address]struct{})


	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		rpcURL = "https://arb1.arbitrum.io/rpc"
	}
	regAddr := os.Getenv("REGISTRY_ADDRESS")
	priv := os.Getenv("PRIVATE_KEY")

	if err := run(context.Background(), rpcURL, regAddr, priv); err != nil {
		log.Fatalf("failed to connect to arbitrum rpc: %v", err)
	}
}
