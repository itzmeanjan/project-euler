package projecteuler

import (
	"math"
	"sync"
)

// holds coefficients (a, b) of quadratic equation
// of form ( n^2 + a*n + b ), and max number of primes
// generated from that equation for continuous values of
// `n`, starting at 0
type coefficient struct {
	a      int
	b      int
	primeC int
}

// checks whether a given number is prime or not
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	target := int(math.Sqrt(float64(n)))
	check := true
	for i := 2; i <= target; i++ {
		if n%i == 0 {
			check = false
			break
		}
	}
	return check
}

// computes maximum number of primes generated from quadratic equation
// of given form, for continuous values of `n`, starting at 0
//
// passes value over channel, to listener go routine
func getPrimeCountFromQuadraticEq(a int, b int, channel chan coefficient) {
	compute := func(n int) int {
		return n*n + a*n + b
	}
	n := 0
	for {
		if !isPrime(compute(n)) {
			break
		}
		n++
	}
	channel <- coefficient{a, b, n}
}

// QuadraticPrimes - Computes product of coefficients of quadratic
// expression, that produces maximum number of primes
// for continuous values of `n`, starting at 0
func QuadraticPrimes() int {
	coeff := coefficient{}
	channel := make(chan coefficient, 1) // communication channel with single value bufferring
	var wg sync.WaitGroup
	wg.Add(1)
	// listener routine
	go func() {
		c := 0
		totalC := 1999 * 2001
		for i := range channel {
			c++
			if i.primeC > coeff.primeC {
				coeff = i
			}
			if c == totalC {
				break
			}
		}
		close(channel)
		wg.Done()
	}()
	for a := -999; a < 1000; a++ {
		for b := -1000; b <= 1000; b++ {
			go getPrimeCountFromQuadraticEq(a, b, channel) // spawning go routines for faster computation, leveraging power of multiple cores of CPU
		}
	}
	wg.Wait()
	return coeff.a * coeff.b
}
