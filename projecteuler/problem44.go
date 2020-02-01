package projecteuler

import "math"

// holds a pair of pentagonal numbers, whose sum
// and absolute difference values are both pentagonal
type pentagonalNumber struct {
	a int
	b int
}

// calculates x-th pentagonal number
func getXthPentagonalNumber(x int) int {
	return x * (3*x - 1) / 2
}

// checks whether a number is pentagonal or not
func isPentagonalNumber(num int) bool {
	check := false
	for i := 1; i*(3*i-1) <= 2*num; i++ {
		tmp := i * (3*i - 1) / 2
		if tmp == num {
			check = true
			break
		}
	}
	return check
}

// tries to find a pair of pentagonal numbers (a, b)
// such that it satisfies following criteria
//
// a+b
// |a-b| - both must be pentagonal
//
// as soon as we find such a pair, we stop searching, and report back
func findPentagonalNumberPair(buffer []int, pos int) pentagonalNumber {
	pair := pentagonalNumber{}
	for i := 0; i < len(buffer); i++ {
		if i == pos {
			continue
		}
		if sum, diff := buffer[i]+buffer[pos], int(math.Abs(float64(buffer[i]-buffer[pos]))); isPentagonalNumber(sum) && isPentagonalNumber(diff) {
			pair.a = buffer[i]
			pair.b = buffer[pos]
			break
		}
	}
	return pair
}

// PentagonalNumbers - Finds a pair of pentagonal numbers,
// addition and absolute difference of them, needs to be pentagonal too
//
// we'll return absolute difference of those two values
func PentagonalNumbers() int {
	pent := pentagonalNumber{}
	buffer := make([]int, 0)
	for i := 1; ; i++ {
		buffer = append(buffer, getXthPentagonalNumber(i))
		if i == 1 {
			continue
		}
		if tmp := findPentagonalNumberPair(buffer, i-1); tmp.a != 0 && tmp.b != 0 {
			pent = tmp
			break
		}
	}
	return int(math.Abs(float64(pent.a - pent.b)))
}
