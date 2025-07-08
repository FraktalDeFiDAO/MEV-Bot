package ethutil

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
)

// ConnectClient connects to an Ethereum RPC endpoint and returns the client.
func ConnectClient(ctx context.Context, url string) (*ethclient.Client, error) {
	return ethclient.DialContext(ctx, url)
}
