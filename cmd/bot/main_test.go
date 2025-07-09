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

	"github.com/FraktalDeFiDAO/MEV-Bot/pkg/market"
	"github.com/FraktalDeFiDAO/MEV-Bot/pkg/watcher"
	"github.com/ethereum/go-ethereum"
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

	bwCalled := false
	ewCalled := false
	var capturedQuery ethereum.FilterQuery
	newBlockWatcher = func(sub watcher.HeaderSubscriber) runner {
		return runnerFunc(func(ctx context.Context) error {
			bwCalled = true
			<-ctx.Done()
			return ctx.Err()
		})
	}
	newEventWatcher = func(sub watcher.LogSubscriber, q ethereum.FilterQuery) runner {
		capturedQuery = q
		return runnerFunc(func(ctx context.Context) error {
			ewCalled = true
			<-ctx.Done()
			return ctx.Err()
		})
	}
	t.Cleanup(func() {
		newBlockWatcher = func(sub watcher.HeaderSubscriber) runner { return watcher.NewBlockWatcher(sub) }
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
	if !bwCalled || !ewCalled {
		t.Fatalf("watchers not called")
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

type stubRegistry struct{ tokens []common.Address }

func (s *stubRegistry) AddPool(a, b, c common.Address, id uint64) (*types.Transaction, error) {
	return new(types.Transaction), nil
}
func (s *stubRegistry) AddToken(t common.Address, d uint8) (*types.Transaction, error) {
	s.tokens = append(s.tokens, t)
	return new(types.Transaction), nil
}

func TestSyncRegistry(t *testing.T) {
	marketStore = &market.Persistent{Market: market.New()}
	addr := common.HexToAddress("0x123")
	marketStore.AddToken(addr)
	stub := &stubRegistry{}
	regClient = stub
	syncRegistry()
	if len(stub.tokens) != 1 || stub.tokens[0] != addr {
		t.Fatalf("registry not updated: %v", stub.tokens)
	}
}
