package watcher

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
)

// HeaderSubscriber wraps SubscribeNewHead so it can be mocked in tests.
type HeaderSubscriber interface {
	SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error)
}

// BlockWatcher listens for new block headers using an Ethereum client.
type BlockWatcher struct {
	sub HeaderSubscriber
}

// NewBlockWatcher creates a watcher using the given subscriber.
func NewBlockWatcher(sub HeaderSubscriber) *BlockWatcher {
	return &BlockWatcher{sub: sub}
}

// Run subscribes to new block headers and logs them until the context is done.
func (bw *BlockWatcher) Run(ctx context.Context) error {
	headers := make(chan *types.Header)
	sub, err := bw.sub.SubscribeNewHead(ctx, headers)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-sub.Err():
			if err != nil {
				return err
			}
		case h := <-headers:
			log.Printf("new block: %d", h.Number.Uint64())
		}
	}
}
