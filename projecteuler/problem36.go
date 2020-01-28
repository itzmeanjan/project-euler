package projecteuler

import "math"

// given a decimal number, converts it
// into its equivalent binary representation
// which is represented using slice
func decToBin(num int) []int {
	size := int(math.Floor(math.Log2(float64(num)))) + 1
	buffer := make([]int, size)
	idx := size - 1
	for num > 0 {
		buffer[idx] = num % 2
		idx--
		num /= 2
	}
	return buffer
}

// given a number, of any base system (2, 10, 8),
// where digits are splitted and stored in a slice,
// checks whether that number is palindrome or not
//
// this one is an improved version of original palindrome
// checker, wrote for decimal numbers
func isPalindromeImp(num []int) bool {
	check := true
	for i, j := 0, len(num)-1; i <= j; {
		if num[i] != num[j] {
			check = false
			break
		}
		i++
		j--
	}
	return check
}

// DoubleBasePalindromes - Calculates sum of all palindromes,
// under 10^6, which are palindrome in both bases (2, 10)
func DoubleBasePalindromes() int {
	sum := 0
	for i := 1; i < 1000000; i++ {
		if IsPalindrome(i) && isPalindromeImp(decToBin(i)) {
			sum += i
		}
	}
	return sum
}
