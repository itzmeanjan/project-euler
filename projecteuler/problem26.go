package projecteuler

// given one fraction in form of numerator and denominator,
// we'll update numerator to certain value so that division can be
// performed
//
// if num = 1 & den = 2; then numerator requires to be updated to 10
// so that division can be performed
// returns updated numerator
func updateNumerator(num int, den int) int {
	for num < den {
		num *= 10
	}
	return num
}

// computes length of recurring cycle in fractional part
// of division
// if division ends i.e. numerator becomes 0,
// we'll say there's no recurring cycle
// but for 1/3 = 0.3333..., we'll say we've a recurring cycle
// of length 1
func getRecurringCycleLength(num int, den int) int {
	length := 0
	buffer := make(map[int]int)
	pos := 0
	buffer[num] = pos
	for num != 0 {
		num = updateNumerator(num, den)
		num %= den
		if _, ok := buffer[num]; ok {
			length = pos - buffer[num] + 1
			break
		}
		pos++
		buffer[num] = pos
	}
	return length
}

// ReciprocalCycle - Finds value of denominator ( < 1000 )
// for which we've max length recurring cycle in fractional part
func ReciprocalCycle() int {
	max := 0
	maxDen := 0
	for i := 2; i < 1000; i++ {
		if tmp := getRecurringCycleLength(1, i); tmp > max {
			max = tmp
			maxDen = i
		}
	}
	return maxDen
}
