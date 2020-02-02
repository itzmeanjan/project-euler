package projecteuler

import "math"

// checks whether a given number can be represented
// as square of some number or not
//
// if given numbers are either 0 or 1, then it's considered to be satisfied
// i.e. 0*0 = 0; 1*1 = 1
func isSquare(num int) bool {
	if num == 0 || num == 1 {
		return true
	}
	check := false
	for i := 2; ; i++ {
		if i*i > num {
			break
		} else if i*i == num {
			check = true
			break
		}
	}
	return check
}

// checks whether a given odd number can be represented
// as sum of one prime number and twice square of a number
//
// for ease of computation, we'll use a buffer i.e. slice of prime numbers
// generated till now
func validateGoldbachConjecture(num int, primes []int) bool {
	check := false
	for i := 0; i < len(primes); i++ {
		if tmp := (num - primes[i]) / 2; tmp >= 0 && isSquare(tmp) {
			check = true
			break
		}
	}
	return check
}

// generates prime numbers after `x`, until we reach `upto`
// and they are stored in same buffer in which pre generated primes
// are kept
func generatePrimesAfterX(x int, upto int, buffer *[]int) {
	for x <= upto {
		x++
		sqrt := int(math.Sqrt(float64(x)))
		check := true
		for i := 0; i < len(*buffer) && (*buffer)[i] <= sqrt; i++ {
			if x%(*buffer)[i] == 0 {
				check = false
				break
			}
		}
		if check {
			*buffer = append(*buffer, x)
		}
	}
}

// GoldbachOtherConjecture - Finds out minimum odd number
// which doesn't satisfy Goldbach's Other Conjecture
func GoldbachOtherConjecture() int {
	odd := 1                            // holds target odd number, to be returned
	buffer := GeneratePrimesUnderX(100) // first generating all primes under 100
	for i := 35; ; i += 2 {             // incrementing by 2, to check only odd numbers
		if tmp := buffer[len(buffer)-1]; i > tmp {
			generatePrimesAfterX(tmp, tmp+(i-tmp+1), &buffer) // generating some next primes, otherwise function fails
			// to check whether next odd number is really satisfying goldbach conjecture or not
		}
		if !validateGoldbachConjecture(i, buffer) { // if doesn't satisfy, we've obtained our target number
			odd = i
			break // then we stop iteration
		}
	}
	return odd
}
