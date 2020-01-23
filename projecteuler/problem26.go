package projecteuler

func doTheyMatch(first, second []int) bool {
	check := true
	for i, j := 0, 0; i < len(first) && j < len(second); {
		if first[i] != second[j] {
			check = false
			break
		}
		i++
		j++
	}
	return check
}

func recurringCyclicPatternHelper(pattern *[]int, start, length int) bool {
	check := true
	for i := start; i+2*length-1 < len(*pattern); i += length {
		if check = doTheyMatch((*pattern)[i:i+length], (*pattern)[i+length:i+2*length]); !check {
			break
		}
	}
	return check
}

func recurringCyclicPattern(pattern *[]int) {
	for i := 0; i <= len(*pattern)/2; i++ {
		length := 1
		for i+length-1 < len(*pattern) {
			recurringCyclicPatternHelper(pattern, i, length)
			length++
		}
	}
}

func fixNumerator(num *int, den int) []int {
	i := 0
	holder := make([]int, 0)
	for *num < den {
		*num *= 10
		if i != 0 {
			holder = append(holder, 0)
		}
		i++
	}
	return holder
}

func divideXTimes(num *int, den int, buffer *[]int, x int) {
	if *num == 0 || x == 0 {
		return
	}
	*buffer = append(*buffer, fixNumerator(num, den)...)
	*buffer = append(*buffer, *num/den)
	*num %= den
	divideXTimes(num, den, buffer, x-1)
}
