package projecteuler

import (
	"fmt"
	"math"
)

type consequtiveNumbers struct {
	a int
	b int
	c int
	d int
}

func (con *consequtiveNumbers) next() {
	con.a = con.b
	con.b = con.c
	con.c = con.d
	con.d++
}

func primeFactorCount(num int, primes []int, factors *map[int]int) {
	for i := 0; i < len(primes); i++ {
		if num < primes[i] {
			break
		} else if num == primes[i] {
			if _, ok := (*factors)[num]; !ok {
				(*factors)[num] = 1
			} else {
				(*factors)[num]++
			}
			break
		} else {
			if num%primes[i] == 0 {
				if _, ok := (*factors)[primes[i]]; !ok {
					(*factors)[primes[i]] = 1
				} else {
					(*factors)[primes[i]]++
				}
				primeFactorCount(num/primes[i], primes, factors)
				break
			}
		}
	}
}

func validatePrimeFactors(num int, factors map[int]int) bool {
	prod := 1
	for k, v := range factors {
		prod *= int(math.Pow(float64(k), float64(v)))
	}
	return num == prod
}

// DistinctPrimeFactors - ...
func DistinctPrimeFactors() int {
	factors := make(map[int]int)
	primes := GeneratePrimesUnderX(17)
	fmt.Println(primes)
	primeFactorCount(17, primes, &factors)
	fmt.Println(factors)
	fmt.Println(validatePrimeFactors(17, factors))
	return 0
}
