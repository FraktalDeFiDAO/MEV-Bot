package main

import (
	"context"
	"os/exec"
	"testing"
	"time"
)

func TestRunInvalidURL(t *testing.T) {
	if err := run(context.Background(), "http://127.0.0.1:0"); err == nil {
		t.Fatal("expected error from invalid RPC URL")
	}
}

func TestRun(t *testing.T) {
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

	if err := run(context.Background(), "http://127.0.0.1:8545"); err != nil {
		t.Fatalf("run failed: %v", err)
	}
}
