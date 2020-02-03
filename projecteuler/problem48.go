package projecteuler

import "math/big"

// given two big numbers, computes sum of them,
// which is no doubt a big number too
func accumulateBig(acc *big.Int, cur *big.Int) *big.Int {
	return acc.Add(acc, cur)
}

// computes a^b, where a & b both can be represented as int64
// but a^b requires big number support library representation
func powerBig(a int, b int) *big.Int {
	c := big.NewInt(int64(1))
	return c.Exp(big.NewInt(int64(a)), big.NewInt(int64(b)), nil)
}

// extracts last `x` digits of a big number, by converting it to string first
func lastXDigitsBig(num big.Int, x int) string {
	tmp := num.String()
	return tmp[len(tmp)-x:]
}

// SelfPowers - Computes a series sum, where each term
// is represented as big number, returns last 10 digits of sum ( big number )
func SelfPowers() string {
	sum := big.NewInt(int64(0))
	for i := 1; i <= 1000; i++ {
		sum = accumulateBig(sum, powerBig(i, i))
	}
	return lastXDigitsBig(*sum, 10)
}
