package projecteuler

// checks whether a given number is abundant number or not
func isAbundantNumber(n int) bool {
	sum := 0
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum > n
}

// generates all abundant numbers under `X`,
// and stores them in slice
func genAbundantNumUnderX(x int) []int {
	cache := make([]int, 0)
	for i := 2; i < x; i++ {
		if isAbundantNumber(i) {
			cache = append(cache, i)
		}
	}
	return cache
}

// determines whether a given number `X`
// can be written as sum of two abundant numbers
// or not
// a cache of all abundant numbers to be passed,
// while invoking this function
func canBeWrittenAsSumOfTwoAbundantNum(x int, cache []int) bool {
	check := false
	for i := 0; i < len(cache); i++ {
		if cache[i] > x {
			break
		}
		for j := i; j < len(cache); j++ {
			tmp := cache[i] + cache[j]
			if tmp == x {
				check = true
				break
			} else if tmp > x {
				break
			}
		}
		if check {
			break
		}
	}
	return check
}

// NonAbundantSum - Calculates sum of all positive integers
// which can't be written as sum of two abundant numbers
func NonAbundantSum() int {
	const limit = 28124
	sum := 0
	tmp := genAbundantNumUnderX(limit)
	for i := 1; i < limit; i++ {
		if !canBeWrittenAsSumOfTwoAbundantNum(i, tmp) {
			sum += i
		}
	}
	return sum
}
