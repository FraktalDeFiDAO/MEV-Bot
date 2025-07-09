package ethutil

import (
	"context"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

// Noncer defines the subset of ethclient.Client used for nonce management.
type Noncer interface {
	PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)
}

// NonceManager manages a monotonically increasing nonce for an account.
type NonceManager struct {
	mu     sync.Mutex
	next   uint64
	addr   common.Address
	client Noncer
}

// NewNonceManager creates a manager using the current pending nonce.
func NewNonceManager(ctx context.Context, c Noncer, addr common.Address) (*NonceManager, error) {
	nonce, err := c.PendingNonceAt(ctx, addr)
	if err != nil {
		return nil, err
	}
	return &NonceManager{next: nonce, addr: addr, client: c}, nil
}

// Next returns the next nonce and increments the counter.
func (n *NonceManager) Next(ctx context.Context) (uint64, error) {
	n.mu.Lock()
	defer n.mu.Unlock()

	if n.next == 0 {
		nonce, err := n.client.PendingNonceAt(ctx, n.addr)
		if err != nil {
			return 0, err
		}
		n.next = nonce
	}

	nonce := n.next
	n.next++
	return nonce, nil
}
