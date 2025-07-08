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
	log.Println("block watcher starting")
	headers := make(chan *types.Header)
	sub, err := bw.sub.SubscribeNewHead(ctx, headers)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	defer log.Println("block watcher stopped")

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

// LogSubscriber wraps SubscribeFilterLogs so it can be mocked in tests.
type LogSubscriber interface {
	SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error)
}

// LogHandler is called for each log received by an EventWatcher.
type LogHandler func(types.Log)

// EventWatcher listens for contract events based on a filter query.
type EventWatcher struct {
	sub     LogSubscriber
	query   ethereum.FilterQuery
	handler LogHandler
}

// NewEventWatcher creates an event watcher for the given query and handler.
func NewEventWatcher(sub LogSubscriber, q ethereum.FilterQuery, h LogHandler) *EventWatcher {
	return &EventWatcher{sub: sub, query: q, handler: h}
}

// Run subscribes to logs and prints their transaction hash until the context ends.
func (ew *EventWatcher) Run(ctx context.Context) error {
	log.Println("event watcher starting")
	logsCh := make(chan types.Log)
	sub, err := ew.sub.SubscribeFilterLogs(ctx, ew.query, logsCh)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()
	defer log.Println("event watcher stopped")

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-sub.Err():
			if err != nil {
				return err
			}
		case l := <-logsCh:
			if ew.handler != nil {
				ew.handler(l)
			} else {
				log.Printf("log tx: %s", l.TxHash.Hex())
			}
		}
	}
}
