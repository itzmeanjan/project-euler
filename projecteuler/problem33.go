package projecteuler

import "math"

// one custom type created for holding a fraction
// i.e. numerator and denominator both seperately
type fraction struct {
	num int
	den int
}

// recursively computes GCD of two given integers
func gcdRec(a, b int) int {
	if b == 0 {
		return a
	}
	return gcdRec(b, a%b)
}

// calculates minimum of two given integers
func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// calculates GCD of two given integers iteratively
func gcd(a, b int) int {
	min, gcd := getMin(a, b), 1
	for i := 1; i <= min; i++ {
		if a%i == 0 && b%i == 0 {
			gcd = i
		}
	}
	return gcd
}

// converts a given number into a slice of digits,
// where MSB is placed at 0th index of slice
func splitDigits(num int) []int {
	size := int(math.Floor(math.Log10(float64(num)))) + 1
	buffer := make([]int, size)
	i := size - 1
	for num > 0 {
		buffer[i] = num % 10
		num /= 10
		i--
	}
	return buffer
}

// checks whether one given fraction is curious or not,
// as per given definition
func isNonTrivial(frac fraction) bool {
	num, den := splitDigits(frac.num), splitDigits(frac.den)
	if num[1] != den[0] {
		return false
	}
	tmp := fraction{num[0], den[1]}
	if tmpGCD, fracGCD := gcd(tmp.num, tmp.den), gcd(frac.num, frac.den); tmp.num/tmpGCD == frac.num/fracGCD && tmp.den/tmpGCD == frac.den/fracGCD {
		return true
	}
	return false
}

// generates all fractions having two digits in both numerator and denominator
// & whose value < 1 && not having `0` as digit in numerator or denominator
func generateFractions(from int, to int) []fraction {
	fractions := make([]fraction, 0)
	for num := from; num < to; num++ {
		if num%10 == 0 {
			continue
		}
		for den := num + 1; den <= to; den++ {
			if den%10 == 0 {
				continue
			}
			fractions = append(fractions, fraction{num, den})
		}
	}
	return fractions
}

// extracts only curious fractions
func getNonTrivialFractions() []fraction {
	buffer := make([]fraction, 0)
	for _, i := range generateFractions(11, 99) {
		if isNonTrivial(i) {
			buffer = append(buffer, i)
		}
	}
	return buffer
}

// multiplies a collection of fractions,
// numerators with numerators & denominators with denominators
// while resulting product is stored in a fraction
func multiplyFractions(frac []fraction) fraction {
	tmp := fraction{1, 1}
	for _, i := range frac {
		tmp.num *= i.num
		tmp.den *= i.den
	}
	return tmp
}

// DigitCancellingFractions - Extracts all curious fractions first, which are then multiplied
// and represented in its lowest common terms
//
// finally denominator of that fraction is returned
func DigitCancellingFractions() int {
	frac := multiplyFractions(getNonTrivialFractions())
	return frac.den / gcd(frac.num, frac.den)
}
