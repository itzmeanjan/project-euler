package projecteuler

import "math"

// given a slice of decimal digits (0-9)
// returns an integer formed by those digits
func numFromDigits(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum = sum*10 + arr[i]
	}
	return sum
}

// given a number, rotates that number by one digit place
// if 197 is given, after single digit rotation it'll be 971
// after that it'll be 719
func rotateNumber(num *[]int) {
	if len(*num) == 1 {
		return
	}
	tmp := make([]int, len(*num))
	copy(tmp, *num)
	for i := 1; i <= len(*num); i++ {
		(*num)[i-1] = tmp[i%len(*num)]
	}
}

// checks whether a given number is circular prime or not
func isCircularPrime(num int) (bool, []int) {
	if !isPrime(num) { // first we check, whether given number is prime or not
		return false, nil // if not, no  need to check for its circular forms
	}
	splitted := splitDigits(num)
	circulated := make([]int, int(math.Floor(math.Log10(float64(num))))+1)
	circulated[0] = num
	check := true
	for i := 1; i < len(circulated); i++ {
		rotateNumber(&splitted)                 // rotating number, for creating next value
		circulated[i] = numFromDigits(splitted) // putting rotated number into buffer
		if !isPrime(circulated[i]) {            // checking whether this form is prime or not, if not
			check = false // we simply quit looping, to reduce CPU cycle usage
			break
		}
	}
	if !check {
		return check, nil
	}
	return check, circulated // if circular prime, all circulated numbers generated to be returned
}

// copies content of slice into a given hash map,
// if that value is not present in hash map
func putIntoBuffer(items []int, buffer *map[int]int) {
	for _, v := range items {
		if _, ok := (*buffer)[v]; !ok {
			(*buffer)[v] = 1
		}
	}
}

// CircularPrimes - Calculates number of circular primes under 10^6
func CircularPrimes() int {
	primeBuffer := make(map[int]int)
	for i := 2; i < 1000000; i++ {
		if _, ok := primeBuffer[i]; !ok {
			if check, extras := isCircularPrime(i); check {
				putIntoBuffer(extras, &primeBuffer)
			}
		}
	}
	return len(primeBuffer)
}
