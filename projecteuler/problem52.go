package projecteuler

// given a number and a slice of multiples
// we'll check whether all products, obtained
// by multiplying `num` with multipliers
// are permutation of same digits
func areMultiplesPermuted(num int, multiplier []int) bool {
	check := true
	for _, v := range multiplier {
		if !permutationOfSameDigits(num, num*v) {
			check = false
			break
		}
	}
	return check
}

// PermutedMultiples - We'll obtain smallest number
// which can generate multiples ( when multiplied with
// each of 2, 3, 4, 5, 6 ), that are permutation of same digits
func PermutedMultiples() int {
	num := 1
	multipliers := []int{2, 3, 4, 5, 6}
	for ; ; num++ {
		if areMultiplesPermuted(num, multipliers) {
			break
		}
	}
	return num
}
