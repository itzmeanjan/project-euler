package projecteuler

// given an integer `num` & a collection of cached primes,
// we'll find out all unique prime factors that number has
//
// lets take an example :
// 14 = 2*7 ( unique prime factor count : 2 )
// 15 = 3*5 ( unique prime factor count : 2 )
// 16 = 2^4 ( unique prime factor count : 1 )
// 17 = 1*17 ( unique prime factor count : 1 )
func primeFactorCount(num int, primes []int, factors *map[int]int) {
	for i := 0; i < len(primes); i++ {
		if num < primes[i] { // as soon as we exceed number under check, we'll quit
			break
		} else if num == primes[i] { // if `num` itself is a prime, then we store it & break out of loop
			if _, ok := (*factors)[num]; !ok {
				(*factors)[num] = 1
			} else {
				(*factors)[num]++
			}
			break
		} else {
			if num%primes[i] == 0 { // if divisible by this prime ( < num ), we'll store it & apply same operation on quotient
				if _, ok := (*factors)[primes[i]]; !ok {
					(*factors)[primes[i]] = 1
				} else {
					(*factors)[primes[i]]++
				}
				primeFactorCount(num/primes[i], primes, factors) // if we've 14, first it'll be discovered that it's divisible by `2`
				break                                            // for further computation we'll push 14/2 ( = 7 ) in , for getting its unique prime count, which is to be cached in same buffer
			}
		}
	}
}

// holds four consequtive integers, and their corresponding
// unique prime factors with repeatation count
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

// obtains next set of four consequtive integers
// by changing their positions
//
// only new element is last element `d`
//
// So its unique prime factor holder
// is reset to empty
//
// 1, 2, 3, 4 will become 2, 3, 4, 5
// so only new number is last one, so for avoiding
// huge recomputation of prime factors, we'll keep
// pre computed factors for a, b, c. But for `d`, it'll be
// set to empty, because it was never computer prior to this point
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

// given a buffer of primes ( cached ), we'll check whether
// current combination of consequtive integers are having `count` number
// of prime factors or not
// all of them must have `count` number of unique prime factors
func (con consequtiveNumbers) primeCountCheck(primes []int, count int) bool {
	primeFactorCount(con.d, primes, &con.factorsD)
	return len(con.factorsA) == count && len(con.factorsA) == len(con.factorsB) && len(con.factorsB) == len(con.factorsC) && len(con.factorsC) == len(con.factorsD)
}

// DistinctPrimeFactors - First four consequtive integers, each having four unique prime factors,
// need to return that first number of combination
func DistinctPrimeFactors() int {
	cons := consequtiveNumbers{1, map[int]int{1: 1}, 2, map[int]int{2: 1}, 3, map[int]int{3: 1}, 4, map[int]int{2: 2}} // beginning with 1, 2, 3, 4 & their corresponding unique prime factor holder
	primes := GeneratePrimesUnderX(100)                                                                                // caching all primes under 100, for avoiding recomputation
	for !cons.primeCountCheck(primes, 4) {                                                                             // until we find four consequtive number combination, each having 4 unique prime factors, keep exploring
		if tmp := primes[len(primes)-1]; cons.d > tmp {
			generatePrimesAfterX(tmp, cons.d+1, &primes) // caching some next primes, in same buffer
		}
		cons.next() // we'll keep exploring next combination
	}
	return cons.a
}
