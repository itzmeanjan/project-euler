package projecteuler

// computes sum of all proper divisors of a given number
// `n` i.e. all numbers <=n/2, which divide `n` evenly.
func getSumOfProperDivisors(n int) int {
	sum := 0
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}

// AmicableNumbers - Computes sum of all
// amicable numbers under 10000
func AmicableNumbers() int {
	sum := 0
	cache := [10000]bool{}
	for i := 1; i < 10000; i++ {
		if cache[i] {
			sum += i
			continue
		}
		tmp := getSumOfProperDivisors(i)
		if tmp != i && i == getSumOfProperDivisors(tmp) {
			cache[i] = true
			cache[tmp] = true
			sum += i
		}
	}
	return sum
}
