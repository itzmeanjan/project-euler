package projecteuler

// PowerfulDigitSum - Computes maximum sum of digits
// of natural number of form a^b where a, b < 100
func PowerfulDigitSum() int {
	max := 0
	for i := 99; i > 0; i-- {
		for j := 99; j > 0; j-- {
			if tmp := digitSum(*powerBig(i, j)); tmp > max {
				max = tmp
			}
		}
	}
	return max
}
