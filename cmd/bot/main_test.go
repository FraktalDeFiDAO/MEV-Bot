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

	"github.com/FraktalDeFiDAO/MEV-Bot/pkg/watcher"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type runnerFunc func(context.Context) error

func (f runnerFunc) Run(ctx context.Context) error { return f(ctx) }

func TestRunInvalidURL(t *testing.T) {
	if err := run(context.Background(), "http://127.0.0.1:0"); err == nil {
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
	newBlockWatcher = func(sub watcher.HeaderSubscriber) runner {
		return runnerFunc(func(ctx context.Context) error {
			bwCalled = true
			<-ctx.Done()
			return ctx.Err()
		})
	}
	newEventWatcher = func(sub watcher.LogSubscriber, q ethereum.FilterQuery) runner {
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

	if err := run(ctx, "http://127.0.0.1:8545"); err != context.Canceled {
		t.Fatalf("run failed: %v", err)
	}
	if !bwCalled || !ewCalled {
		t.Fatalf("watchers not called")
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
