package projecteuler

import (
	"math/big"
)

// finds sum of digits, given a big number ( big.Int )
func digitSum(bn big.Int) int {
	sum := 0
	for _, i := range bn.String() {
		sum += (int(i) - 48)
	}
	return sum
}

// calculates factorial of any number as big.Int,
// so it's possible to calculate factorial of any number
func factorial(n int) big.Int {
	res := big.NewInt(int64(1))
	for i := n; i > 1; i-- {
		res = res.Mul(res, big.NewInt(int64(i)))
	}
	return *res
}

// FactorialDigitSum - Finds sum of digits of factorial of any given number `n`
func FactorialDigitSum(n int) int {
	return digitSum(factorial(n))
}
