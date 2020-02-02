package projecteuler

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

type consequtiveNumbers struct {
	a        int
	factorsA map[int]int
	b        int
	factorsB map[int]int
	c        int
	factorsC map[int]int
	d        int
	factorsD map[int]int
}

func (con *consequtiveNumbers) next() {
	con.a = con.b
	con.factorsA = con.factorsB
	con.b = con.c
	con.factorsB = con.factorsC
	con.c = con.d
	con.factorsC = con.factorsD
	con.d++
	con.factorsD = make(map[int]int)
}

func (con consequtiveNumbers) primeCountCheck(primes []int, count int) bool {
	primeFactorCount(con.d, primes, &con.factorsD)
	return len(con.factorsA) == count && len(con.factorsA) == len(con.factorsB) && len(con.factorsB) == len(con.factorsC) && len(con.factorsC) == len(con.factorsD)
}

// DistinctPrimeFactors - ...
func DistinctPrimeFactors() int {
	cons := consequtiveNumbers{1, map[int]int{1: 1}, 2, map[int]int{2: 1}, 3, map[int]int{3: 1}, 4, map[int]int{2: 2}}
	primes := GeneratePrimesUnderX(100)
	for !cons.primeCountCheck(primes, 4) {
		if tmp := primes[len(primes)-1]; cons.d > tmp {
			generatePrimesAfterX(tmp, cons.d+1, &primes)
		}
		cons.next()
	}
	return cons.a
}
