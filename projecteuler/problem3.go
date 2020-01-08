package projecteuler

import "math"

// GeneratePrimesUnderX - Generates all prime numbers under one given `X`,
// and returns a slice of them
//
// using dynamic programming strategy, to perform lesser number of checkings
func GeneratePrimesUnderX(x int) []int {
	primeArr := []int{2}
	for i := 3; i <= x; i++ {
		sqrt := int(math.Sqrt(float64(i)))
		check := true
		for j := 0; j < len(primeArr) && primeArr[j] <= sqrt; j++ {
			if i%primeArr[j] == 0 {
				check = false
				break
			}
		}
		if check {
			primeArr = append(primeArr, i)
		}
	}
	return primeArr
}

// GetLargestPrimeFactor - First calculates square root of given number,
// and find out all primes which are under or equals to that sqrt value
//
// Now we'll simply iterate over that prime holder slice, from last to first,
// i.e. from higher value prime to lower value prime, cause finally, we need
// to find out maximum prime factor of `num`. That'll allow us to perform lesser
// number of checkings.
func GetLargestPrimeFactor(num int) int {
	largest := 0
	primeArr := GeneratePrimesUnderX(int(math.Sqrt(float64(num))))
	for i := len(primeArr) - 1; i >= 0; i-- {
		if num%primeArr[i] == 0 {
			largest = primeArr[i]
			break
		}
	}
	return largest
}
