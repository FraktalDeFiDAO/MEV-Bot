package main

import (
	"context"
	"log"

	"github.com/FraktalDeFiDAO/MEV-Bot/pkg/ethutil"
)

// connectClient abstracts ethutil.ConnectClient for testability.
var connectClient = ethutil.ConnectClient

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
	return nil
}

func main() {
	rpcURL := "https://arb1.arbitrum.io/rpc"

	if err := run(context.Background(), rpcURL); err != nil {
		log.Fatalf("failed to connect to arbitrum rpc: %v", err)
	}

	// TODO: subscribe to dex events, transactions and sequencer feed
}
