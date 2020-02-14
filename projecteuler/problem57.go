package projecteuler

import "math/big"

// represents √2 in form of rational number
type squareRootOf2 struct {
	num *big.Int
	den *big.Int
}

// finds next expanded rational number
// which can represent √2
//
// in first iteration we represent it using 3/2
// in next expansion, we get 7/5
// next to it, 17/12
// we can find a clear pattern in calculating
// next expanded form
// num(i) = num(i-1)+ den(i-1)*2
// den(i) = num(i-1) + den(i-1)
// p/q = num(i)/ den(i)
func (sqrt *squareRootOf2) next() {
	num, den := sqrt.num, sqrt.den
	sqrt.num = big.NewInt(int64(0)).Add(num, big.NewInt(int64(1)).Mul(den, big.NewInt(int64(2))))
	sqrt.den = big.NewInt(int64(0)).Add(num, den)
}

// checks whether a given fraction
// is having more digits in numerator than in denominator or not
func (sqrt squareRootOf2) isValid() bool {
	return len(sqrt.num.String()) > len(sqrt.den.String())
}

// SquareRootConvergents - Computes number of fractions ( obtained by expanding √2 )
// satisfying criteria i.e. number of digits in numerator more than digits in denominator
func SquareRootConvergents() int {
	count := 0
	sqrt := squareRootOf2{big.NewInt(int64(3)), big.NewInt(int64(2))}
	for i := 0; i < 1000; i++ {
		if sqrt.isValid() {
			count++
		}
		sqrt.next()
	}
	return count
}
