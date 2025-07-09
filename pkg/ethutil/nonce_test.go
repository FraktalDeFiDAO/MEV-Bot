package ethutil

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

type nonceClient struct {
	nonce uint64
}

func (n *nonceClient) PendingNonceAt(ctx context.Context, addr common.Address) (uint64, error) {
	return n.nonce, nil
}

func TestNonceManager(t *testing.T) {
	ctx := context.Background()
	c := &nonceClient{nonce: 7}
	nm, err := NewNonceManager(ctx, c, common.HexToAddress("0x1"))
	if err != nil {
		t.Fatal(err)
	}
	n1, err := nm.Next(ctx)
	if err != nil || n1 != 7 {
		t.Fatalf("unexpected nonce %d err %v", n1, err)
	}
	n2, _ := nm.Next(ctx)
	if n2 != 8 {
		t.Fatalf("expected 8 got %d", n2)
	}
}
