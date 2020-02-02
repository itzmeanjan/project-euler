package projecteuler

import (
	"math"
)

// swaps two integers by their address
func swapNum(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

// given a sequence of integers, we'll start reverse
// elemenets from given start index to end index
func reverseNum(arr []int, start int, end int) {
	for i := 0; i < (end-start+1)/2; i++ {
		swapNum(&arr[start+i], &arr[end-i])
	}
}

// given a sequence of integers, next lexicographic permutation to
// be generated by Narayana Pandit's Algorithm, 14th Centuary
func generateNextPermNum(arr []int) {
	last := len(arr) - 1 // last index of set
	if last < 1 {        // if we've only one element in set, nothing to permute
		return
	}
	i := last - 1
	for i >= 0 && !(arr[i] < arr[i+1]) { // finds out highest index ( i ) in set such that arr[i] < arr[i+1]
		i--
	}
	if i < 0 { // if it's lesser than 0, we've generated last possible sequence, reverting back to initial combination
		reverseNum(arr, 0, last)
	} else {
		j := last
		for j > i+1 && !(arr[j] > arr[i]) { // finding out index ( j ) in set such that arr[i] < arr[j]
			j--
		}
		swapNum(&arr[i], &arr[j])  // swapping arr[i] & arr[j]
		reverseNum(arr, i+1, last) // reversing all elements of set, starting from index `i+1` to `last`
	}
	// and now we've next lexicographic permutation / reverted back to initial combination
}

// given a sequence of numbers, we'll group them
// into sub strings of specified length, starting from given index
// and putting repeation of `repeat` number of elements
// from previous sub string
func groupify(num []int, c int, start int, repeat int) []int {
	buffer := make([]int, 0)
	for i := start; i <= len(num)-c; i += (c - repeat) {
		buffer = append(buffer, numFromDigits(num[i:i+c], 10))
	}
	return buffer
}

// generates first `x` prime numbers, and stores them in a slice
func getFirstXPrimes(x int) []int {
	buffer := make([]int, x)
	buffer[0] = 2
	num := 3
	for i := 1; i < x; {
		check := true
		sqrt := int(math.Sqrt(float64(num)))
		for j := 0; j < i && buffer[j] <= sqrt; j++ {
			if num%buffer[j] == 0 {
				check = false
				break
			}
		}
		if check {
			buffer[i] = num
			i++
		}
		num += 2
	}
	return buffer
}

// checks whether first element of subs gets evenly
// divided by first prime; second element of subs gets evenly
// divided by second prime; ... and keep going this way.
//
// breaks out of loop, as soon as a check fails
//
// we're trying to check sub string divisibility
// using primes of corresponding position
func checkDivisibility(subs []int, primes []int) bool {
	check := true
	for i, v := range primes {
		if subs[i]%v != 0 {
			check = false
			break
		}
	}
	return check
}

// SubStringDivisibility - Finds sum of all 0-9 pandigital numbers, which satisfies
// given property
//
// starting check from 1023456789, because that's smallest 0-9 pandigital number
// which can be formed, while starting with 1 i.e. starting a number with 0,
// doesn't really make sense
//
// and we'll keep generating lexicographic permutations, which are no doubt 0-9 pandigital
// all we've to do is to groupify them as per specified rule and them their corresponding
// divisibility, until first digit of permutation becomes 0
//
// sum of those pandigital numbers, satisfying this criteria, to be returned
func SubStringDivisibility() int {
	sum := 0
	primes := getFirstXPrimes(7)
	init := splitDigits(1023456789)
	for init[0] != 0 {
		if checkDivisibility(groupify(init, 3, 1, 2), primes) {
			sum += numFromDigits(init, 10)
		}
		generateNextPermNum(init)
	}
	return sum
}