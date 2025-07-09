package market

import (
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestLoadAndSave(t *testing.T) {
	f := "test_market.db"
	defer os.Remove(f)
	pm := LoadFromFile(f)
	if len(pm.List()) != 0 {
		t.Fatal("expected empty")
	}
	pm.AddPool(common.HexToAddress("0x1"), common.HexToAddress("0xa"), common.HexToAddress("0xb"))
	pm.AddToken(common.HexToAddress("0x2"))
	pm2 := LoadFromFile(f)
	if !pm2.Has(common.HexToAddress("0x1")) || !pm2.HasToken(common.HexToAddress("0x2")) {
		t.Fatal("persist failed")
	}
	if p, ok := pm2.GetPool(common.HexToAddress("0x1")); !ok || p.Token0 != common.HexToAddress("0xa") {
		t.Fatalf("pool info missing: %+v", p)
	}
}
