package projecteuler

func areConcatenatedNumbersPrime(num1 int, num2 int) bool {
	return isPrime(joinNumbers(num1, num2)) && isPrime(joinNumbers(num2, num1))
}
