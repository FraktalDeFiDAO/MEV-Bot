package main

import (
	"bytes"
	"context"
	"log"
	"math/big"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/FraktalDeFiDAO/MEV-Bot/pkg/arb"
	"github.com/FraktalDeFiDAO/MEV-Bot/pkg/market"
	"github.com/FraktalDeFiDAO/MEV-Bot/pkg/registry"
	"github.com/FraktalDeFiDAO/MEV-Bot/pkg/watcher"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type runnerFunc func(context.Context) error

func (f runnerFunc) Run(ctx context.Context) error { return f(ctx) }

func TestRunInvalidURL(t *testing.T) {
	if err := run(context.Background(), "http://127.0.0.1:0", "", ""); err == nil {
		t.Fatal("expected error from invalid RPC URL")
	}
}

func TestRun(t *testing.T) {
	if _, err := exec.LookPath("anvil"); err != nil {
		t.Skip("anvil not installed")
	}

	ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, "anvil", "--port", "8545", "--chain-id", "31337")
	cmd.Stdout = nil
	cmd.Stderr = nil
	if err := cmd.Start(); err != nil {
		t.Fatalf("failed to start anvil: %v", err)
	}
	t.Cleanup(func() {
		cancel()
		cmd.Process.Kill()
		cmd.Wait()
	})

	// wait briefly for anvil to start
	time.Sleep(100 * time.Millisecond)

	ewCalled := false
	var capturedQuery ethereum.FilterQuery
	newEventWatcher = func(sub watcher.LogSubscriber, q ethereum.FilterQuery) runner {
		capturedQuery = q
		return runnerFunc(func(ctx context.Context) error {
			ewCalled = true
			<-ctx.Done()
			return ctx.Err()
		})
	}
	t.Cleanup(func() {
		newEventWatcher = func(sub watcher.LogSubscriber, q ethereum.FilterQuery) runner {
			return watcher.NewEventWatcher(sub, q, profitLogHandler)
		}
	})

	go func() {
		// give run time to connect and start watchers before
		// canceling the context so the test can exit cleanly
		time.Sleep(500 * time.Millisecond)
		cancel()
	}()

	if err := run(ctx, "http://127.0.0.1:8545", "", ""); err != context.Canceled {
		t.Fatalf("run failed: %v", err)
	}
	if !ewCalled {
		t.Fatalf("event watcher not called")
	}
	if len(capturedQuery.Topics) == 0 || len(capturedQuery.Topics[0]) != 2 {
		t.Fatalf("unexpected query topics: %v", capturedQuery.Topics)
	}
}

func TestProfitLogHandler(t *testing.T) {
	buf := &bytes.Buffer{}
	log.SetOutput(buf)
	log.SetFlags(0)
	defer func() {
		log.SetOutput(os.Stderr)
		log.SetFlags(log.LstdFlags)
	}()

	// profit event
	data, _ := tradeABI.Pack("TradeExecuted", big.NewInt(1), big.NewInt(2))
	profitLogHandler(types.Log{Topics: []common.Hash{tradeEventID}, Data: data, TxHash: common.HexToHash("0x1")})
	if !strings.Contains(buf.String(), "profit event") {
		t.Fatal("profit event not logged")
	}
	buf.Reset()

	// sync event
	data, _ = syncABI.Events["Sync"].Inputs.Pack(big.NewInt(5), big.NewInt(10))
	profitLogHandler(types.Log{Topics: []common.Hash{syncEventID}, Data: data, TxHash: common.HexToHash("0x2"), Address: common.HexToAddress("0x3")})
	if !strings.Contains(buf.String(), "price update") {
		t.Fatal("price update not logged")
	}
}

func TestPairLogHandler(t *testing.T) {
	buf := &bytes.Buffer{}
	log.SetOutput(buf)
	log.SetFlags(0)
	defer func() {
		log.SetOutput(os.Stderr)
		log.SetFlags(log.LstdFlags)
	}()

	pairData, _ := pairABI.Events["PairCreated"].Inputs.Pack(common.Address{}, common.Address{}, common.HexToAddress("0x111"), big.NewInt(0))
	pairLogHandler(types.Log{Topics: []common.Hash{pairEventID}, Data: pairData})
	if !strings.Contains(buf.String(), "pair created") {
		t.Fatal("pair not logged")
	}
	buf.Reset()

	poolData, _ := poolABI.Events["PoolCreated"].Inputs.Pack(common.Address{}, common.Address{}, big.NewInt(3000), big.NewInt(60), common.HexToAddress("0x222"))
	pairLogHandler(types.Log{Topics: []common.Hash{poolEventID}, Data: poolData})
	if !strings.Contains(buf.String(), "pool created") {
		t.Fatalf("pool not logged: %q", buf.String())
	}
}

func TestPairLogHandlerRegistry(t *testing.T) {
	marketStore = &market.Persistent{Market: market.New()}
	stub := &stubRegistry{}
	regClient = stub
	knownTokens = make(map[common.Address]struct{})
	knownPools = make(map[common.Address]struct{})

	t0 := common.HexToAddress("0xabc")
	t1 := common.HexToAddress("0xdef")
	pair := common.HexToAddress("0x123")
	data, _ := pairABI.Events["PairCreated"].Inputs.Pack(t0, t1, pair, big.NewInt(0))
	pairLogHandler(types.Log{Topics: []common.Hash{pairEventID, common.BytesToHash(t0.Bytes()), common.BytesToHash(t1.Bytes())}, Data: data})

	if len(stub.tokens) != 2 || len(stub.pools) != 1 {
		t.Fatalf("unexpected registry calls: %v %v", stub.tokens, stub.pools)
	}

	// sending the same event again should not trigger new txs
	pairLogHandler(types.Log{Topics: []common.Hash{pairEventID, common.BytesToHash(t0.Bytes()), common.BytesToHash(t1.Bytes())}, Data: data})
	if len(stub.tokens) != 2 || len(stub.pools) != 1 {
		t.Fatalf("duplicate registry calls: %v %v", stub.tokens, stub.pools)
	}
}

type stubRegistry struct {
	tokens []common.Address
	pools  [][3]common.Address
}

type stubExec struct{ called bool }

func (s *stubExec) Execute(opts *bind.TransactOpts, a, b common.Address, maxIn, step *big.Int) (*types.Transaction, error) {
	s.called = true
	return types.NewTx(&types.LegacyTx{Nonce: 0}), nil
}

type stubNonce struct{ n uint64 }

func (s *stubNonce) Next(context.Context) (uint64, error) {
	s.n++
	return s.n - 1, nil
}

func (s *stubRegistry) AddPool(a, b, c common.Address, id uint64) (*types.Transaction, error) {
	s.pools = append(s.pools, [3]common.Address{a, b, c})
	tx := types.NewTx(&types.LegacyTx{Nonce: 0})
	return tx, nil
}
func (s *stubRegistry) AddToken(t common.Address, d uint8) (*types.Transaction, error) {
	s.tokens = append(s.tokens, t)
	tx := types.NewTx(&types.LegacyTx{Nonce: 0})
	return tx, nil
}

func (s *stubRegistry) WaitMined(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	return &types.Receipt{Status: types.ReceiptStatusSuccessful}, nil
}

func (s *stubRegistry) Tokens(context.Context) ([]common.Address, error) {
	return s.tokens, nil
}

func (s *stubRegistry) Pools(context.Context) ([]common.Address, error) {
	var res []common.Address
	for _, p := range s.pools {
		res = append(res, p[0])
	}
	return res, nil
}

func (s *stubRegistry) PoolInfo(ctx context.Context, a common.Address) (registry.PoolInfo, error) {
	if len(s.pools) == 0 {
		return registry.PoolInfo{}, nil
	}
	p := s.pools[0]
	return registry.PoolInfo{Token0: p[1], Token1: p[2], ExchangeID: big.NewInt(0), Enabled: true}, nil
}

func TestSyncRegistry(t *testing.T) {
	marketStore = &market.Persistent{Market: market.New()}
	addr := common.HexToAddress("0x123")
	marketStore.AddToken(addr)
	marketStore.AddToken(common.HexToAddress("0x456"))
	marketStore.AddPool(common.HexToAddress("0x1"), addr, common.HexToAddress("0x456"))
	stub := &stubRegistry{}
	regClient = stub
	syncRegistry()
	if len(stub.tokens) != 2 {
		t.Fatalf("registry tokens not updated: %v", stub.tokens)
	}
	if len(stub.pools) != 1 {
		t.Fatalf("registry pools not updated: %v", stub.pools)
	}
}

func TestOpportunityHandler(t *testing.T) {
	arbExec = &stubExec{}
	execAuth = &bind.TransactOpts{}
	nonceMgr = &stubNonce{}
	arbMon = arb.NewMonitor([][2]common.Address{{common.HexToAddress("0x1"), common.HexToAddress("0x2")}}, 10, 1)
	arbMon.SetHandler(opportunityHandler)
	arbMon.Update(common.HexToAddress("0x1"), big.NewInt(1000), big.NewInt(1000))
	arbMon.Update(common.HexToAddress("0x2"), big.NewInt(1200), big.NewInt(800))
	if !arbExec.(*stubExec).called {
		t.Fatalf("executor not called")
	}
}
