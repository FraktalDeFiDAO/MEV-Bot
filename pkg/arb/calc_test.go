package arb

import (
	"math/big"
	"testing"
)

func TestFindBestInput(t *testing.T) {
	rA0 := big.NewInt(1000)
	rB0 := big.NewInt(1000)
	rA1 := big.NewInt(1200)
	rB1 := big.NewInt(800)
	bestIn, profit := FindBestInput(rA0, rB0, rA1, rB1, 500, 1)
	if bestIn.Cmp(big.NewInt(0)) <= 0 {
		t.Fatalf("expected positive input")
	}
	if profit.Cmp(big.NewInt(0)) <= 0 {
		t.Fatalf("expected positive profit")
	}
}
