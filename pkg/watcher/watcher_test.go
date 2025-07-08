package watcher

import (
	"context"
	"errors"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
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

type errLogSubscriber struct{ err error }

func (e errLogSubscriber) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	return nil, e.err
}

type goodLogSubscriber struct{}

func (goodLogSubscriber) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	sub := stubSub{errCh: make(chan error)}
	go func() {
		ch <- types.Log{TxHash: common.HexToHash("0x1")}
		time.Sleep(10 * time.Millisecond)
		close(sub.errCh)
	}()
	return sub, nil
}

type errAfterSubscriber struct{}

func (errAfterSubscriber) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
	sub := stubSub{errCh: make(chan error, 1)}
	go func() {
		ch <- &types.Header{Number: big.NewInt(1)}
		sub.errCh <- errors.New("boom")
	}()
	return sub, nil
}

type errAfterLogSubscriber struct{}

func (errAfterLogSubscriber) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	sub := stubSub{errCh: make(chan error, 1)}
	go func() {
		ch <- types.Log{TxHash: common.HexToHash("0x2")}
		sub.errCh <- errors.New("boom")
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

func TestEventWatcherSubscribeError(t *testing.T) {
	ew := NewEventWatcher(errLogSubscriber{err: errors.New("boom")}, ethereum.FilterQuery{})
	if err := ew.Run(context.Background()); err == nil {
		t.Fatal("expected error")
	}
}

func TestEventWatcherRun(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ew := NewEventWatcher(goodLogSubscriber{}, ethereum.FilterQuery{})
	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()
	if err := ew.Run(ctx); err != context.Canceled {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestBlockWatcherSubscriptionError(t *testing.T) {
	bw := NewBlockWatcher(errAfterSubscriber{})
	if err := bw.Run(context.Background()); err == nil || err.Error() != "boom" {
		t.Fatalf("expected boom, got %v", err)
	}
}

func TestEventWatcherSubscriptionError(t *testing.T) {
	ew := NewEventWatcher(errAfterLogSubscriber{}, ethereum.FilterQuery{})
	if err := ew.Run(context.Background()); err == nil || err.Error() != "boom" {
		t.Fatalf("expected boom, got %v", err)
	}
}
