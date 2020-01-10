package projecteuler

import "math"

// GetXthPrime - Calculates `X`-th prime number
// where, starting one is 2.
//
// we're trying to reduce number of checks
// by testing only odd numbers ( cause even numbers
// will always be composite )
//
// while checking whether a certain number is prime
// or not, we'll only test their divisibility by primes
// which are lesser than squareroot of number under lens,
// that'll allow us to perform lesser checks.
//
// Using, every composite number has prime factors.
func GetXthPrime(x int) int {
	primeArr := make([]int, 1, x)
	primeArr[0] = 2
	num := 3
	for len(primeArr) != x {
		sqrt := int(math.Sqrt(float64(num)))
		check := true
		for j := 0; j < len(primeArr) && primeArr[j] <= sqrt; j++ {
			if num%primeArr[j] == 0 {
				check = false
				break
			}
		}
		if check {
			primeArr = append(primeArr, num)
		}
		num += 2
	}
	return primeArr[x-1]
}
