package projecteuler

import "math"

// checks whether a given multiplicand, multiplier, product identity is
// pandigital or not
func isIdentityPandigital(multiplicand int, multiplier int, product int) bool {
	buffer := make(map[int]int)                        // stores occurance count of digits
	doesDigitOccurMoreThanOnes := func(num int) bool { // checks whether any digit occurs more than ones or not ( in given number )
		check := false
		for num > 0 {
			if digit := num % 10; digit == 0 {
				check = true
				break
			} else {
				if _, ok := buffer[digit]; ok {
					check = true
					break
				} else {
					buffer[digit] = 1
				}
			}
			num /= 10
		}
		return check
	}
	checkOverallDigitPresence := func(low int, high int) bool { // checks whether 1-9 all digits are present exactly 1's or not
		check := true
		for i := low; i <= high; i++ {
			if v, ok := buffer[i]; !ok || v != 1 {
				check = false
				break
			}
		}
		return check
	}
	return !doesDigitOccurMoreThanOnes(multiplicand) && !doesDigitOccurMoreThanOnes(multiplier) && !doesDigitOccurMoreThanOnes(product) && checkOverallDigitPresence(1, 9)
}

// given a map of {int: int} type, where key is holding
// items to be summed; and values holding corresponding occurance count
//
// we'll calculate sum of those keys present in map
func getSum(buffer map[int]int) int {
	sum := 0
	for i := range buffer {
		sum += i
	}
	return sum
}

// PandigitalProducts - Finds sum of all those pandigital identities,
// for which we've unique products
func PandigitalProducts() int {
	productsExplored := make(map[int]int)
	digitC := func(num int) int {
		return int(math.Floor(math.Log10(float64(num)))) + 1
	}
	for i := 2; i < 10000; i++ {
		for j := i + 1; j < 10000; j++ {
			prod := i * j
			if c := digitC(i) + digitC(j) + digitC(prod); c == 9 {
				if _, ok := productsExplored[prod]; !ok && isIdentityPandigital(i, j, prod) {
					productsExplored[prod] = 1
				}
			} else if c > 9 {
				break
			} else {
				continue
			}
		}
	}
	return getSum(productsExplored)
}
