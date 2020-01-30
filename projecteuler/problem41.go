package projecteuler

import "sync"

// calculates first decimal number
// that requires `n` digits for representation
//
// i.e. first 2-digit number 10
func firstNDigitNumber(n int) int {
	num := 1
	for n > 1 {
		num *= 10
		n--
	}
	return num
}

// max decimal number that requires `n` digits
// for presentation
//
// i.e. 99 is max number that requires 2 digits
func lastNDigitNumber(n int) int {
	return firstNDigitNumber(n+1) - 1
}

// finds maximum number that require `n` digits
// which is also pandigital ( from 1 to n) and prime
func largestNDigitPandigitalPrime(n int, channel chan int) {
	prime := 0
	end := firstNDigitNumber(n)
	for i := lastNDigitNumber(n); i >= end; i -= 2 {
		if isPandigital(splitDigits(i), 1, n) && isPrime(i) {
			prime = i
			break
		}
	}
	channel <- prime
}

// PandigitalPrime - Computes maximum number that is pandigital and prime
func PandigitalPrime() int {
	channel := make(chan int, 1)
	maxPrime := 0
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		c := 0
		for i := range channel {
			if i > maxPrime {
				maxPrime = i
			}
			c++
			if c == 8 {
				break
			}
		}
		close(channel)
		wg.Done()
	}()
	for i := 9; i > 1; i-- {
		go largestNDigitPandigitalPrime(i, channel)
	}
	wg.Wait()
	return maxPrime
}
