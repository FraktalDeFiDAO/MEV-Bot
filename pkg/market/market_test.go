package market

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestMarket(t *testing.T) {
	m := New()
	addr := common.HexToAddress("0x1")
	if m.Has(addr) {
		t.Fatal("should not have addr yet")
	}
	m.Add(addr)
	if !m.Has(addr) {
		t.Fatal("expected addr to be present")
	}
	list := m.List()
	if len(list) != 1 || list[0] != addr {
		t.Fatalf("unexpected list: %v", list)
	}
}
