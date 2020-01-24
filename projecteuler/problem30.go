package projecteuler

import "math"

// calculates sum of digits ( of a given number ) with x-th power
func sumOfDigitsPowerX(num int, x int) int {
	sum := 0
	tmp := num
	for num > 0 {
		sum += int(math.Pow(float64(num%10), float64(x)))
		if sum > tmp {
			break
		}
		num /= 10
	}
	return sum
}

// verifies whether `num` equals to sum of digits of ( `num` ) elevated to x-th power
func verifySum(num int, x int) bool {
	return sumOfDigitsPowerX(num, x) == num
}

// DigitFifthPowers - Calculates sum of all numbers that can be
// written as sum of fifth power of their digits
func DigitFifthPowers() int {
	sum := 0
	for i := 10; i < 1000000; i++ {
		if verifySum(i, 5) {
			sum += i
		}
	}
	return sum
}
