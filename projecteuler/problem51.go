package projecteuler

import (
	"math"
)

// given a set of numbers, extracts only those which are prime
func getPrimesOnly(num []int) []int {
	buffer := make([]int, 0)
	for _, v := range num {
		if isPrime(v) {
			buffer = append(buffer, v)
		}
	}
	return buffer
}

// given a number as a slice i.e. digits are
// splitted & put as an element of slice
// e.g. 123 will be kept as []int{1, 2, 3}
// and another slice of positions to be supplied,
// length of which is <= length of `num` slice
//
// all we've to do is replace all those indices
// given in `pos` slice, with digits from 0-9
// so if num = []int{1, 2, 3} & pos = []int{0}
// then generated numbers []int{123, 223,
// 323, 423, 523, 623, 723, 823, 923}, but 023 = 23 not be
// considered, because it's not 3 digit anymore
func replaceXDigitsBy0_9(pos []int, num []int) []int {
	buffer := make([]int, 0)
	for i := 0; i <= 9; i++ {
		tmp := getCopy(num)
		for _, j := range pos {
			tmp[j] = i
		}
		if v := numFromDigits(tmp, 10); int(math.Floor(math.Log10(float64(v))))+1 == len(num) {
			buffer = append(buffer, v)
		}
	}
	return buffer
}

// given a set of numbers, all we've to do
// is to choose all those instances of given length
// from that set
// if []int{1, 2, 3} is given, and asked to choose all unordered subsets
// of length 2, then []int{[]int{1, 2}, []int{1, 3}, []int{2, 3}} - choices we can make
func getPossibleChoices(frm []int, chooseX int) [][]int {
	choices := binomialCoeff(len(frm), chooseX)
	buffer := make([][]int, choices)
	init := make([]int, chooseX)
	for i := 0; i < chooseX; i++ {
		init[i] = i
	}
	for i := 0; i < choices; i++ {
		buffer[i] = getCopy(init)
		chooseNext(frm, init)
	}
	return buffer
}

// given a prime number ( `num` ), and `x` as number of primes
// we want to generate by replacing some indices of `num`
// with 0-9, we'll find starting value of that series
// i.e. minimum value of prime series with length `x`
func getSmallestPrimeFromXPrimeFamily(num int, x int) int {
	base := make([]int, int(math.Floor(math.Log10(float64(num))))+1)
	for i := range base {
		base[i] = i
	}
	splittedN := splitDigits(num)
	initP := 0
	for i := 1; i < len(base); i++ {
		for _, j := range getPossibleChoices(base, i) {
			if primes := getPrimesOnly(replaceXDigitsBy0_9(j, getCopy(splittedN))); len(primes) == x {
				initP = primes[0]
				break
			}
		}
		if initP != 0 {
			break
		}
	}
	return initP // returns 0, if found nothing
}

// PrimeDigitReplacements - Replaces digits of prime numbers
// with 0-9 and generates prime numbers of same length as of
// that primary prime number. We'll try to find out such a
// series of primes, of length `x`, and return minimum
// number of that series
// x = 8, here
func PrimeDigitReplacements(x int) int {
	smallestP := 0
	for i := 2; ; i++ {
		for _, p := range generateAllXDigitsPrimes(i) {
			if v := getSmallestPrimeFromXPrimeFamily(p, x); v != 0 {
				smallestP = v
				break
			}
		}
		if smallestP != 0 {
			break
		}
	}
	return smallestP
}
