package projecteuler

// given a slice of digits, holding a decimal number,
// this function checks whether number is pandigital
// from digit `frm` through digit `to`
// so number of digits need to ( to-frm+1 )
func isPandigital(num []int, frm int, to int) bool {
	if len(num) != to-frm+1 {
		return false
	}
	check := true
	buffer := make(map[int]int)
	for _, v := range num {
		if _, ok := buffer[v]; ok || !(v >= frm && v <= to) {
			check = false
			break
		}
		buffer[v] = 1
	}
	return check
}

// given a slice of numbers, we'll simply split each number
// into its digits, and concatenate them, while maintaining their order
//
// so , {1, 12, 123, 1234} will become {1, 1, 2, 1, 2, 3, 1, 2, 3, 4}
func concatenateNum(num ...int) []int {
	buffer := make([]int, 0)
	for _, v := range num {
		buffer = append(buffer, splitDigits(v)...)
	}
	return buffer

}

// given a multiplicand & a range of multipliers
// it'll return a slice holding products in order ( ascending in multiplier )
func generateMultiples(multiplicand int, multF int, multT int) []int {
	buffer := make([]int, multT-multF+1)
	c := 0
	for i := multF; i <= multT; i++ {
		buffer[c] = multiplicand * i
		c++
	}
	return buffer
}

// PandigitalMultiples - Finds maximum 1 through 9 pandigital
// number that can be obtained by multiplying `X`, with {1, 2, .. , n}, where n > 1
// and concatenating them
func PandigitalMultiples() int {
	max := 0
	for i := 2; i <= 9876; i++ {
		for j := 2; ; j++ {
			num := concatenateNum(generateMultiples(i, 1, j)...) // generated concatenated number
			if ln := len(num); ln > 9 {
				break
			} else if ln < 9 {
				continue
			}
			if !isPandigital(num, 1, 9) { // checking whether pandigital or not
				break
			}
			if tmp := numFromDigits(num, 10); max < tmp { // checking whether it has exceeded `max` or not
				max = tmp
				break
			}
		}
	}
	return max
}
