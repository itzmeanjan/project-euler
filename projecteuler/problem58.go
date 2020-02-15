package projecteuler

// given prime number count & total number count,
// we need to compute precentage ratio of
// prime numbers to all numbers, as integer
func primeToAllRatio(primeC int, allC int) int {
	return primeC * 100 / allC
}

// given a slice of numbers, we'll find out number of primes
func getPrimeCount(buffer []int) int {
	c := 0
	for _, v := range buffer {
		if isPrime(v) {
			c++
		}
	}
	return c
}

// we're asked to generate only diagonal elements
// of a clockwise rotating spiral, where last element generated
// in previous iteration of spiral is given, along with current side length of spiral
// side length of spiral is same as width of square matrix, currently under consideration
// returns four elements generated in this run, present on four corners i.e. on diagonals
// of square matrix
func generateDiagonalElementsOfClockWiseSpiral(lastV *int, sideL int, initSideL int) []int {
	buffer := make([]int, 4)
	incrBy := sideL - initSideL
	for i := 0; i < 4; i++ {
		*lastV += incrBy
		buffer[i] = *lastV
	}
	return buffer
}

// SpiralPrimes - Finds out at which side length value of clockwise spiral
// we'll get a prime ratio of value < 10%
// avoid prime checking recomputation, by precaching count of already
// checked primes
func SpiralPrimes() int {
	lastV, sideL, primeC, totalC := 1, 3, 0, 1
	for ; ; sideL += 2 {
		primeC += getPrimeCount(generateDiagonalElementsOfClockWiseSpiral(&lastV, sideL, 1))
		totalC += 4
		if primeToAllRatio(primeC, totalC) < 10 {
			break
		}
	}
	return sideL
}
