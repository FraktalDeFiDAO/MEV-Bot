package market

import (
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestLoadAndSave(t *testing.T) {
	f := "test_market.json"
	defer os.Remove(f)
	pm := LoadFromFile(f)
	if len(pm.List()) != 0 {
		t.Fatal("expected empty")
	}
	pm.Add(common.HexToAddress("0x1"))
	pm.AddToken(common.HexToAddress("0x2"))
	pm2 := LoadFromFile(f)
	if !pm2.Has(common.HexToAddress("0x1")) || !pm2.HasToken(common.HexToAddress("0x2")) {
		t.Fatal("persist failed")
	}
}
