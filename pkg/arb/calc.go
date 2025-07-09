package arb

import "math/big"

const (
	feeNumerator   = 997
	feeDenominator = 1000
)

func getAmountOut(amountIn, reserveIn, reserveOut *big.Int) *big.Int {
	amountInWithFee := new(big.Int).Mul(amountIn, big.NewInt(feeNumerator))
	numerator := new(big.Int).Mul(amountInWithFee, reserveOut)
	denominator := new(big.Int).Add(new(big.Int).Mul(reserveIn, big.NewInt(feeDenominator)), amountInWithFee)
	return new(big.Int).Div(numerator, denominator)
}

func GetProfit(amountIn, rA0, rB0, rA1, rB1 *big.Int) *big.Int {
	out1 := getAmountOut(amountIn, rA0, rB0)
	out2 := getAmountOut(out1, rB1, rA1)
	if out2.Cmp(amountIn) > 0 {
		return new(big.Int).Sub(out2, amountIn)
	}
	return big.NewInt(0)
}

func FindBestInput(rA0, rB0, rA1, rB1 *big.Int, maxIn, step int64) (bestIn, bestProfit *big.Int) {
	bestIn = big.NewInt(0)
	bestProfit = big.NewInt(0)
	for i := step; i <= maxIn; i += step {
		amt := big.NewInt(i)
		p := GetProfit(amt, rA0, rB0, rA1, rB1)
		if p.Cmp(bestProfit) > 0 {
			bestProfit = p
			bestIn = new(big.Int).Set(amt)
		}
	}
	return
}
