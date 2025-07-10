package watcher

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
)

// HeaderSubscriber wraps SubscribeNewHead so it can be mocked in tests.
type HeaderSubscriber interface {
	SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error)
}

// BlockWatcher listens for new block headers using an Ethereum client.
type BlockWatcher struct {
	sub     HeaderSubscriber
	verbose bool
}

// NewBlockWatcher creates a watcher using the given subscriber.
func NewBlockWatcher(sub HeaderSubscriber) *BlockWatcher {
	return &BlockWatcher{sub: sub, verbose: os.Getenv("DEBUG") != ""}
}

// Run subscribes to new block headers and logs them until the context is done.
func (bw *BlockWatcher) Run(ctx context.Context) error {
	log.Println("block watcher starting")
	headers := make(chan *types.Header)
	defer log.Println("block watcher stopped")

	for {
		sub, err := bw.sub.SubscribeNewHead(ctx, headers)
		if err != nil {
			return err
		}

	inner:
		for {
			select {
			case <-ctx.Done():
				sub.Unsubscribe()
				return ctx.Err()
			case err := <-sub.Err():
				if err != nil {
					log.Printf("block subscription error: %v", err)
				}
				sub.Unsubscribe()
				time.Sleep(time.Second)
				break inner
			case h := <-headers:
				if bw.verbose {
					log.Printf("new block: %d", h.Number.Uint64())
				}
			}
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
	defer log.Println("event watcher stopped")

	for {
		sub, err := ew.sub.SubscribeFilterLogs(ctx, ew.query, logsCh)
		if err != nil {
			return err
		}

	inner:
		for {
			select {
			case <-ctx.Done():
				sub.Unsubscribe()
				return ctx.Err()
			case err := <-sub.Err():
				if err != nil {
					log.Printf("log subscription error: %v", err)
				}
				sub.Unsubscribe()
				time.Sleep(time.Second)
				break inner
			case l := <-logsCh:
				if ew.handler != nil {
					ew.handler(l)
				} else {
					log.Printf("log tx: %s", l.TxHash.Hex())
				}
			}
		}
	}
}
