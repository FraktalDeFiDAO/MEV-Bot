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
	if m.HasToken(addr) {
		t.Fatal("should not have token yet")
	}
	m.Add(addr)
	m.AddToken(addr)
	if !m.Has(addr) || !m.HasToken(addr) {
		t.Fatal("expected addr to be present")
	}
	list := m.List()
	if len(list) != 1 || list[0] != addr {
		t.Fatalf("unexpected list: %v", list)
	}
	tlist := m.ListTokens()
	if len(tlist) != 1 || tlist[0] != addr {
		t.Fatalf("unexpected token list: %v", tlist)
	}
}
