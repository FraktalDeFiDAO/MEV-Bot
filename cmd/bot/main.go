package main

import (
	"context"
	"log"
	"os"
	"github.com/FraktalDeFiDAO/MEV-Bot/pkg/ethutil"
	"github.com/FraktalDeFiDAO/MEV-Bot/pkg/watcher"
	"github.com/ethereum/go-ethereum"
)

type runner interface{ Run(context.Context) error }

// connectClient abstracts ethutil.ConnectClient for testability.
var (
	connectClient   = ethutil.ConnectClient
	newBlockWatcher = func(sub watcher.HeaderSubscriber) runner { return watcher.NewBlockWatcher(sub) }
	newEventWatcher = func(sub watcher.LogSubscriber, q ethereum.FilterQuery) runner { return watcher.NewEventWatcher(sub, q) }
)

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

func main() {
	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		rpcURL = "https://arb1.arbitrum.io/rpc"
	}

	if err := run(context.Background(), rpcURL); err != nil {
		log.Fatalf("failed to connect to arbitrum rpc: %v", err)
	}
}
