package projecteuler

import "math/big"

// uses golang big number support package
// for computing exponent
func powerOf(base int64, pow int64) string {
	res := new(big.Int)
	res = res.Exp(big.NewInt(base), big.NewInt(pow), nil)
	return res.String()
}

// PowerDigitSum - Calculates sum of digits of calculated exponent
// for given `base` and `pow` values
func PowerDigitSum(base int, pow int) int {
	sum := 0
	for _, i := range powerOf(int64(base), int64(pow)) {
		sum += int(i - 48)
	}
	return sum
}
