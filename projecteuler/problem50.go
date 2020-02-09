package projecteuler

// holds longest consecutive prime sequence
// length & sum of those elements
type consecutivePrimes struct {
	length int
	sum    int
}

// given a slice of primes, finds sum of
// primes, given by `from` - `to` index ( including both )
func consecutivePrimeSum(primes []int, from int, to int) int {
	sum := 0
	for i := from; i <= to; i++ {
		sum += primes[i]
	}
	return sum
}

// given a slice of primes, find sum of
// longest prime sequence, which is having sum below specified
// value
func findLongestSumOfConsecutivePrimes(primes []int, below int) int {
	cons := consecutivePrimes{}
	holdIt := consecutivePrimeSum(primes, 0, len(primes)-1) // at very beginning, we start with computing sum of all primes present in slice
	for i := 0; i < len(primes); i++ {                      // outer loop index points at beginning of current sequence under inspection
		sum := holdIt                                     // we've whole sequence sum ready here, which is to be updated, in each iteration of inner loop
		for j := len(primes) - 1; j >= cons.length; j-- { // in each iteration we'll keep decreasing inner loop ( i.e. right most index ) by one
			if j-i+1 <= cons.length { // sequence length is lesser than current max sequence length
				break
			}
			if j < len(primes)-1 {
				sum -= primes[j+1]
			}
			if sum < below && isPrime(sum) { // checking whether sum is prime or not
				cons.length = j - i + 1
				cons.sum = sum
			}
		}
		holdIt -= primes[i] // computing sum of series, from updated left index to last index ( len(primes) - 1 )
	}
	return cons.sum
}

// ConsecutivePrimeSum - Finds sum of longest consecutive prime series
// which is also a prime, and lesser than 10^6
func ConsecutivePrimeSum() int {
	return findLongestSumOfConsecutivePrimes(GeneratePrimesUnderX(1000000), 1000000)
}
