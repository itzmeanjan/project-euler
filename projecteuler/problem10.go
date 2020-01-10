package projecteuler

// SumOfPrimes - Calculates sum of all primes
// under a given number `x`
func SumOfPrimes(x int) int {
	sum := 0
	for _, i := range GeneratePrimesUnderX(x) {
		sum += i
	}
	return sum
}
