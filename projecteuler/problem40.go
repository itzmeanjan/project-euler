package projecteuler

import "math"

// custom type for holding current position in an irrational number
// being generated
type irrationalFraction struct {
	cur int
	n   int
}

// finds number of digits is a given
// number `x`
func digitCount(x int) int {
	return int(math.Floor(math.Log10(float64(x)))) + 1
}

// we'll not store all 10^6 digits in a slice, so
// that later we can compute required equation value
// rather we'll choose one tricky way, we'll only keep 1-5 digits ( at max )
// at a time, and find stopDigit positions i.e. 1, 10, 100, 1000, 10000, 100000, 1000000
// and keep digit at corresponding position in a hash map, which will be
// later used for generating result
func generateIrrationalFraction(upto int) map[int]int {
	frac := irrationalFraction{1, 1}
	cache := make(map[int]int)
	stopAt := 10
	xthDigit := func(x int) {
		i := frac.n - digitCount(frac.cur) + 1
		for _, v := range splitDigits(frac.cur) {
			if x == i {
				cache[x] = v
				break
			}
			i++
		}
	}
	for frac.n <= upto {
		if stopAt == frac.n || (stopAt >= frac.n-digitCount(frac.cur) && stopAt < frac.n) {
			xthDigit(stopAt)
			stopAt *= 10
		}
		frac.cur++
		frac.n += digitCount(frac.cur)
	}
	return cache
}

// ChampernownesConstant - Finds product of digits
// at given positions of a irrational fraction, generated in
// previous function
func ChampernownesConstant() int {
	prod := 1
	for _, v := range generateIrrationalFraction(1000000) {
		prod *= v
	}
	return prod
}
