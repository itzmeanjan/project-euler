package projecteuler

import (
	"math/big"
)

// computes possible ways to choose k elements
// from n elements ( k <= n )
//
// for optimizing computation (because we're using big number library
// which is expensive ), we'll use same pascal triangle
// for choosing mF to mT elements from n elements i.e. for choosing
// 1 to 3 elements from 10 elements
func binomialCoeffBig(n int, mF int, mT int) []*big.Int {
	buffer := make([][]*big.Int, n+1)
	for i := range buffer {
		buffer[i] = make([]*big.Int, n+1)
	}
	for i := 0; i < n+1; i++ {
		for j := 0; j < n+1; j++ {
			buffer[i][j] = big.NewInt(int64(0))
		}
	}
	for i := range buffer {
		buffer[i][i] = big.NewInt(int64(1))
	}
	for i := 1; i < n+1; i++ {
		buffer[i][0] = big.NewInt(int64(1))
	}
	for i := 2; i < n+1; i++ {
		for j := 1; j <= mT; j++ {
			buffer[i][j] = buffer[i][j].Add(buffer[i-1][j], buffer[i-1][j-1])
		}
	}
	return buffer[n][mF : mT+1]
}

// determines how many of ways will be more than 10^6
// when asked to choose 1 - x element subsets from x element set
func getCountOfChoicesGT10_6ForSelectFromX(x int) int {
	count := 0
	target := big.NewInt(int64(1000000))
	for _, v := range binomialCoeffBig(x, 1, x) {
		if v.Cmp(target) == 1 {
			count++
		}
	}
	return count
}

// CombinatorialSelection - Computes how many of nCr will be > 10^6
// where 1 <= n <= 100 ( & no doubt r<=n in each iteration )
func CombinatorialSelection() int {
	count := 0
	for i := 23; i < 101; i++ {
		count += getCountOfChoicesGT10_6ForSelectFromX(i)
	}
	return count
}
