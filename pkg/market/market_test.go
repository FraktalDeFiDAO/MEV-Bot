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
	token0 := common.HexToAddress("0xa")
	token1 := common.HexToAddress("0xb")
	m.AddPool(addr, token0, token1)
	m.AddToken(token0)
	m.AddToken(token1)
	if !m.Has(addr) || !m.HasToken(token0) || !m.HasToken(token1) {
		t.Fatal("expected entries to be present")
	}
	if p, ok := m.GetPool(addr); !ok || p.Token0 != token0 || p.Token1 != token1 {
		t.Fatalf("unexpected pool data: %+v", p)
	}
	list := m.List()
	if len(list) != 1 || list[0] != addr {
		t.Fatalf("unexpected list: %v", list)
	}
	if len(m.ListPools()) != 1 {
		t.Fatal("expected pool list size 1")
	}
}
