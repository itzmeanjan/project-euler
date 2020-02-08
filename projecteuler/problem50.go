package projecteuler

type consecutivePrimes struct {
	length int
	sum    int
}

func consecutivePrimeSum(primes []int, from int, to int) int {
	sum := 0
	for i := from; i <= to; i++ {
		sum += primes[i]
	}
	return sum
}

func findLongestSumOfConsecutivePrimes(primes []int, below int) int {
	cons := consecutivePrimes{}
	holdIt := consecutivePrimeSum(primes, 0, len(primes)-1)
	for i := 0; i < len(primes); i++ {
		sum := holdIt
		for j := len(primes) - 1; j >= cons.length; j-- {
			if j-i+1 <= cons.length {
				break
			}
			if j < len(primes)-1 {
				sum -= primes[j+1]
			}
			if sum < below && isPrime(sum) {
				cons.length = j - i + 1
				cons.sum = sum
			}
		}
		holdIt -= primes[i]
	}
	return cons.sum
}

// ConsecutivePrimeSum - ...
func ConsecutivePrimeSum() int {
	return findLongestSumOfConsecutivePrimes(GeneratePrimesUnderX(1000000), 1000000)
}
