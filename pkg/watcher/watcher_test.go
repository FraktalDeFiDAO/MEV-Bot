package watcher

import (
	"context"
	"errors"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
)

type errSubscriber struct{ err error }

func (e errSubscriber) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
	return nil, e.err
}

type stubSub struct{ errCh chan error }

func (s stubSub) Unsubscribe()      {}
func (s stubSub) Err() <-chan error { return s.errCh }

type goodSubscriber struct{}

func (goodSubscriber) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
	sub := stubSub{errCh: make(chan error)}
	go func() {
		ch <- &types.Header{Number: big.NewInt(1)}
		time.Sleep(10 * time.Millisecond)
		close(sub.errCh)
	}()
	return sub, nil
}

func TestBlockWatcherSubscribeError(t *testing.T) {
	bw := NewBlockWatcher(errSubscriber{err: errors.New("boom")})
	if err := bw.Run(context.Background()); err == nil {
		t.Fatal("expected error")
	}
}

func TestBlockWatcherRun(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	bw := NewBlockWatcher(goodSubscriber{})
	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()
	if err := bw.Run(ctx); err != context.Canceled {
		t.Fatalf("unexpected error: %v", err)
	}
}
