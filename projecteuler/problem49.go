package projecteuler

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

func chooseNext(arr []int, cur []int) []int {
	s := cur[len(cur)-1]
	pos := len(cur) - 1
	for !(s < len(arr)-len(cur)+pos) {
		pos--
		if pos < 0 {
			break
		}
		s = cur[pos]
	}
	if pos == -1 {
		return func() []int {
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

func joinNumbers(num ...int) int {
	buffer := make([]int, 0)
	for _, v := range num {
		buffer = append(buffer, splitDigits(v)...)
	}
	return numFromDigits(buffer, 10)
}

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

// PrimePermutations - ...
func PrimePermutations() int {
	primes := generateAllXDigitsPrimes(4)
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
		chooseNext(primes, cur)
	}
	return targetV
}
