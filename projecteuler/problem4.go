package projecteuler

// IsPalindrome - Checks whether a given number
// is Palindrome or not
//
// If a number reads same from ltr and rtl
// then it's a Palindrome Number
func IsPalindrome(num int) bool {
	tmp := num
	rev := 0
	for tmp > 0 {
		rev = rev*10 + tmp%10
		tmp /= 10
	}
	return rev == num
}

// GetLargestPalindrome - Calculates largest possible palindrome number
// from product of two numbers under `X`
//
// As per problem statement, this function will be invoked
// with 1000, cause we need to calculate largest possible palindrome number
// from multiplication of two three digit number ( i.e. < 1000 )
func GetLargestPalindrome(x int) int {
	largestPalim := 1
	for i := x - 1; i > 0; i-- {
		for j := i; j > 0; j-- {
			tmp := i * j
			if tmp < largestPalim {
				break
			}
			if IsPalindrome(tmp) && tmp > largestPalim {
				largestPalim = tmp
			}
		}
	}
	return largestPalim
}
