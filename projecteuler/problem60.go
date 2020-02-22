package projecteuler

import "fmt"

func populatePrimeBufferWithXElements(primeBuffer *[]bool, x int) {
	start := len(*primeBuffer)
	tmp := make([]bool, x)
	*primeBuffer = append(*primeBuffer, tmp...)
	for i := start; i < len(*primeBuffer); i++ {
		if isPrime(i) {
			(*primeBuffer)[i] = true
		}
	}
}

func primeCheckUsingBuffer(num int, primeBuffer *[]bool) bool {
	diff := num - len(*primeBuffer)
	if diff < 0 {
		return (*primeBuffer)[num]
	}
	populatePrimeBufferWithXElements(primeBuffer, diff+1)
	return (*primeBuffer)[num]
}

func areConcatenatedNumbersPrime(num1 int, num2 int, primeBuffer *[]bool) bool {
	return primeCheckUsingBuffer(joinNumbers(num1, num2), primeBuffer) && primeCheckUsingBuffer(joinNumbers(num2, num1), primeBuffer)
}

func isPrimePairSet(primes []int, primeBuffer *[]bool) bool {
	check := true
	for _, v := range getPossibleChoices(primes, 2) {
		if !areConcatenatedNumbersPrime(primes[v[0]], primes[v[1]], primeBuffer) {
			check = false
			break
		}
	}
	return check
}

// PrimePairSets - ...
func PrimePairSets() int {
	primes := make([]bool, 0)
	fmt.Println(isPrimePairSet([]int{3, 7, 109, 673}, &primes))
	return 0
}
