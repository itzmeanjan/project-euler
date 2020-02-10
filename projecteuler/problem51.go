package projecteuler

import "math"

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
	return initP
}

// PrimeDigitReplacements - ...
func PrimeDigitReplacements() int {
	smallestP := 0
	for i := 2; ; i++ {
		for _, p := range generateAllXDigitsPrimes(i) {
			if v := getSmallestPrimeFromXPrimeFamily(p, 8); v != 0 {
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
