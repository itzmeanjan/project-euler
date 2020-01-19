package projecteuler

import "math/big"

// ThousandDigitFibNumber - Computes index of first fib term, having 1000 digits
func ThousandDigitFibNumber() int {
	// calculates number of digits present in a big.Int number
	digitCount := func(x *big.Int) int {
		return len(x.String())
	}
	// initializing cache, for sake of dynamic programming
	cache := make([]*big.Int, 2)
	cache[0] = big.NewInt(int64(0))        // fib(0) = 0
	cache[1] = big.NewInt(int64(1))        // fib(1) = 1
	next := 2                              // next index, for which fib term to be generated
	for digitCount(cache[next-1]) < 1000 { // loop keeps running, until we've obtained first fib term with 1000 digits
		tmp := big.NewInt(int64(0))
		tmp = tmp.Add(cache[next-1], cache[next-2]) // computed next term
		cache = append(cache, tmp)                  // appended next term to cache
		next++                                      // next fib index, term to be generated
	}
	return next - 1 // index of first fib term, having 1000 digits
}
