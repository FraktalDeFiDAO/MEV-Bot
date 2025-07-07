package main

import (
    "context"
    "log"

    "github.com/ethereum/go-ethereum/ethclient"
)

// Entry point for the MEV bot. Connects to an Arbitrum node and listens for events.


// ConnectClient connects to an Ethereum RPC and returns the client.
func ConnectClient(ctx context.Context, url string) (*ethclient.Client, error) {
    return ethclient.DialContext(ctx, url)
}

// Entry point for the MEV bot. Connects to an Arbitrum node and listens for events.

func main() {
    // replace with your Arbitrum RPC endpoint
    rpcURL := "https://arb1.arbitrum.io/rpc"

    client, err := ConnectClient(context.Background(), rpcURL)


    if err != nil {
        log.Fatalf("failed to connect to arbitrum rpc: %v", err)
    }
    defer client.Close()

    log.Println("connected to arbitrum", rpcURL)

    // TODO: subscribe to dex events, transactions and sequencer feed
}

