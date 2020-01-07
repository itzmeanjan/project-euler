package projecteuler

// FindSumOfAllMultiplesOf3or5BelowX - Calculates sum of
// all multiples of 3 or 5 below a given number `X`
func FindSumOfAllMultiplesOf3or5BelowX(x int) int {
	sum := 0
	for i := 3; i < x; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}
