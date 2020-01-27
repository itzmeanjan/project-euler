package projecteuler

// first converts given number into a slice
// of digits, then each digit is updated with its corresponding
// factorial values, which is precomputed in a map, simply looking up
// from that buffer will be enough
func digitFactorial(num int, buffer *map[int]int) []int {
	tmp := splitDigits(num)
	for i := range tmp {
		tmp[i] = (*buffer)[tmp[i]]
	}
	return tmp
}

// computes sum of items of a slice i.e.
// returns sum of computed slice by previous function
func digitFactorialSum(arr []int) int {
	sum := 0
	for _, i := range arr {
		sum += i
	}
	return sum
}

// checks whether a given number is curious number
// or not
func isCuriousNumber(num int, buffer *map[int]int) bool {
	return num == digitFactorialSum(digitFactorial(num, buffer))
}

// creates a hash map, holding factorial of all digits
// of decimal number system, which will be used
// for looking up factorial value of certain digit, while
// computing digit factorial sum
func makeFactorialBuffer() map[int]int {
	factorial := func(n int) int {
		prod := 1
		for i := n; i > 1; i-- {
			prod *= i
		}
		return prod
	}
	buffer := make(map[int]int)
	for i := 0; i < 10; i++ {
		buffer[i] = factorial(i)
	}
	return buffer
}

// DigitFactorial - Computes sum of all curious numbers
// NOTE: there are only two curious numbers 145 & 40225
func DigitFactorial() int {
	buffer := makeFactorialBuffer()
	sum := 0
	for i := 10; i < 41000; i++ {
		if isCuriousNumber(i, &buffer) {
			sum += i
		}
	}
	return sum
}
