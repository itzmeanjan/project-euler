package projecteuler

// SquareOfSum - Calculates square of sum of all natural
// numbers upto `X`
func SquareOfSum(x int) int {
	tmp := x * (x + 1) / 2
	return tmp * tmp
}

// SumOfSquares - Calculates sum of squares
// of all natural numbers upto `X`
func SumOfSquares(x int) int {
	sum := 1
	for i := 2; i <= x; i++ {
		sum += i * i
	}
	return sum
}

// SumSquareDiff - Calculates difference of result
// of previous two functions ( when invoked with `x` )
func SumSquareDiff(x int) int {
	return SquareOfSum(x) - SumOfSquares(x)
}
