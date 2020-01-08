package projecteuler

// GetSumOfEvenValuedFibonacciTermsUnderX - Calculates sum of all even fibonacci terms
// not crossing given value `X`
//
// Used dynamic programming style, for calculating fibonacci terms, otherwise,
// recursion will make computation very straightforward but slow ( and no doubt very expensive )
func GetSumOfEvenValuedFibonacciTermsUnderX(x int) int {
	sum := 2
	fibArr := []int{1, 2}
	i := 2
	for {
		tmp := fibArr[i-1] + fibArr[i-2]
		if tmp >= x {
			break
		}
		fibArr = append(fibArr, tmp)
		if tmp%2 == 0 {
			sum += tmp
		}
		i++
	}
	return sum
}
