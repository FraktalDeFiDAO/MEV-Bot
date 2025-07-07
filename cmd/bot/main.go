package main

import (
    "context"
    "log"
<<<<<<< HEAD
    "github.com/ethereum/go-ethereum/ethclient"
)

// Entry point for the MEV bot. Connects to an Arbitrum node and listens for events.
=======

    "github.com/ethereum/go-ethereum/ethclient"
)

// ConnectClient connects to an Ethereum RPC and returns the client.
func ConnectClient(ctx context.Context, url string) (*ethclient.Client, error) {
    return ethclient.DialContext(ctx, url)
}

// Entry point for the MEV bot. Connects to an Arbitrum node and listens for events.

>>>>>>> s9mh77-codex/create-foundry.toml-for-solidity-testing
func main() {
    // replace with your Arbitrum RPC endpoint
    rpcURL := "https://arb1.arbitrum.io/rpc"

<<<<<<< HEAD
    client, err := ethclient.DialContext(context.Background(), rpcURL)
=======
    client, err := ConnectClient(context.Background(), rpcURL)
>>>>>>> s9mh77-codex/create-foundry.toml-for-solidity-testing
    if err != nil {
        log.Fatalf("failed to connect to arbitrum rpc: %v", err)
    }
    defer client.Close()

    log.Println("connected to arbitrum", rpcURL)

    // TODO: subscribe to dex events, transactions and sequencer feed
}

