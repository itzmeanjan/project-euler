package projecteuler

// GetXthTriangularNumber - ...
func GetXthTriangularNumber(x int) int {
	return x * (x + 1) / 2
}

// GetFactorCount - ...
func GetFactorCount(x int) int {
	if x == 1 {
		return 1
	}
	count := 2
	for i := 2; i <= x/2; i++ {
		if x%i == 0 {
			count++
		}
	}
	return count
}

// HighlyDivisibleTriangularNumber - ...
func HighlyDivisibleTriangularNumber() int {
	triNum := 1
	factorC := 1
	i := 1
	for factorC <= 500 {
		triNum = GetXthTriangularNumber(i)
		if tmp := GetFactorCount(triNum); tmp > factorC {
			factorC = tmp
		}
		i++
	}
	return triNum
}
