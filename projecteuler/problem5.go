package projecteuler

// IsDivisible - Checks whether given number
// `X` is divisible by all numbers [`frm` - `to`],
//
// returns boolean result
func IsDivisible(x int, frm int, to int) bool {
	check := true
	for i := frm; i <= to; i++ {
		if x%i != 0 {
			check = false
			break
		}
	}
	return check
}

// GetSmallestMultiple - Returns smallest number
// which is divisible by all numbers starting from
// 1 to 20
//
// Here, I'll not check all numbers, rather for faster
// computation I'd prefer to skip some of them
// As we're asked to find out smallest number divisible by (1..20),
// I'd start from 10, and keep incrementing by 10 ( after each iteration ),
// which will reduce #-of computational steps to a great extent
func GetSmallestMultiple() int {
	num := 10
	for {
		if IsDivisible(num, 1, 20) {
			break
		}
		num += 10
	}
	return num
}
