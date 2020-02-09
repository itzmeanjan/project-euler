package projecteuler

// finds number of ways we can choose
// m number of elements from n elements, where n >= m
//
// using pascal triangle for computing so
func binomialCoeff(n int, m int) int {
	buffer := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		buffer[i] = make([]int, n+1)
		buffer[i][0] = 1
	}
	for i := 0; i < n+1; i++ {
		buffer[i][i] = 1
	}
	for i := 2; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			buffer[i][j] = buffer[i-1][j] + buffer[i-1][j-1]
		}
	}
	return buffer[n][m]
}

// given a decimal number, computes all unique
// digits along with their frequency of presence in that number
func getUniqueDigits(num int) map[int]int {
	digits := make(map[int]int)
	for num > 0 {
		if _, ok := digits[num%10]; ok {
			digits[num%10]++
		} else {
			digits[num%10] = 1
		}
		num /= 10
	}
	return digits
}

// given two numbers, it'll check whether it's possible
// to write those two numbers using permutation of same digits
//
// if 123 & 321 are given, then they are permutation of same digits {1, 2, 3}
// if 969 & 699 is given, then they are permutations of same digits {6, 9}
// but 102 & 211 is not permutation of same digits, cause 0 isn't present in second number
func permutationOfSameDigits(n1 int, n2 int) bool {
	check := true
	n1Map := getUniqueDigits(n1)
	n2Map := getUniqueDigits(n2)
	for i := range n1Map {
		if _, ok := n2Map[i]; !ok {
			check = false
			break
		}
	}
	for i := range n2Map {
		if _, ok := n1Map[i]; !ok {
			check = false
			break
		}
	}
	return check
}

// given a sequence of three primes, it checks
// whether they are arithemetic progression series
// of them or not
// they all need to be permutation of same digits
// which is checked using property of trasitivity
//
// i.e. (a & b) & (b & c) are permutation of same digits
// then it's ensured that (a & c) are permutation of those same digits
func isValidIncreasingSequence(arr []int) bool {
	check := true
	diff := 0
	for i := 0; i < len(arr)-1; i++ {
		if i == 0 {
			diff = arr[i+1] - arr[i]
		}
		if i > 0 && diff != (arr[i+1]-arr[i]) {
			check = false
			break
		}
		if !permutationOfSameDigits(arr[i], arr[i+1]) {
			check = false
			break
		}
	}
	return check
}

// returns a copy of given slice
func getCopy(arr []int) []int {
	tmp := make([]int, len(arr))
	copy(tmp, arr)
	return tmp
}

// given two sets, where first one holds some values, and
// second one holds some select indices of those values,
// our objective is to replace all those indices with value present
// in corresponding index in first set
//
// returns updated set, holding values at given indices
func fillUp(values []int, indices []int) []int {
	for i, v := range indices {
		indices[i] = values[v]
	}
	return indices
}

// given a set of elements to choose from
// and a current chosen sequence of certain length
// we need to find out which will be next sequence to be chosen
// from that set
// if last sequence has been reached, then returns initial one
func chooseNext(arr []int, cur []int) []int {
	pos := len(cur) - 1
	s := cur[pos]
	for !(s < len(arr)-len(cur)+pos) {
		pos--
		if pos < 0 {
			break
		}
		s = cur[pos]
	}
	if pos == -1 { // when we've reached last sequence, which could be chosen from that set
		return func() []int { // computes initial sequence
			for i := range cur {
				cur[i] = i
			}
			return cur
		}()
	}
	cur[pos]++
	for i := pos + 1; i < len(cur); i++ {
		cur[i] = cur[i-1] + 1
	}
	return cur
}

// computes all primes of `x` digit length
func generateAllXDigitsPrimes(x int) []int {
	buffer := make([]int, 0)
	f1 := firstNDigitNumber(x) + 1
	l1 := lastNDigitNumber(x) + 1
	for i := f1; i < l1; i += 2 {
		if isPrime(i) {
			buffer = append(buffer, i)
		}
	}
	return buffer
}

// given a varible number of integers,
// we'll compute a number where all digits get concatinated
// i.e. 123 & 199 gets converted into `123199`
func joinNumbers(num ...int) int {
	buffer := make([]int, 0)
	for _, v := range num {
		buffer = append(buffer, splitDigits(v)...)
	}
	return numFromDigits(buffer, 10)
}

// checks whether content of two slice are same or not
func areTheySame(sl1 []int, sl2 []int) bool {
	check := true
	for i := range sl1 {
		if sl1[i] != sl2[i] {
			check = false
			break
		}
	}
	return check
}

// PrimePermutations - Finds that 12-digit number formed by
// concatinating all digits of three four-digit lengthy primes,
// which forms an AP series ( arithmetic progression )
// and need to be other than 148748178147
func PrimePermutations() int {
	primes := generateAllXDigitsPrimes(4) // computed all primes having four digits
	cur := []int{0, 1, 2}
	notIt := []int{1487, 4817, 8147}
	targetV := 0
	for i := binomialCoeff(len(primes), 3); i > 0; i-- {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		tmp = fillUp(primes, tmp)
		if isValidIncreasingSequence(tmp) && !areTheySame(tmp, notIt) {
			targetV = joinNumbers(tmp...)
			break
		}
		chooseNext(primes, cur) // selects three primes from series
	}
	return targetV
}
