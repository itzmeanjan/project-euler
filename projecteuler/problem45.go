package projecteuler

// calculates x-th hexagonal number
func getXthHexagonalNumber(x int) int {
	return x * (2*x - 1)
}

// checks whether a given number is hexagonal or not
func isHexagonalNumber(num int) bool {
	check := false
	for i := 1; i*(2*i-1) <= num; i++ {
		if i*(2*i-1) == num {
			check = true
			break
		}
	}
	return check
}

// TriangularPentagonalHexagonal - Finds second triangular number
// which is also pentagonal & hexagonal
func TriangularPentagonalHexagonal() int {
	num := 0
	for i := 286; ; i++ {
		if tmp := GetXthTriangularNumber(i); isPentagonalNumber(tmp) && isHexagonalNumber(tmp) {
			num = tmp
			break
		}
	}
	return num
}
