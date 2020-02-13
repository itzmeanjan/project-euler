package projecteuler

import (
	"math/big"
	"sync"
)

// given a big integer, reverses it
// e.g. if 123 is given, 321 to be returned
// doesn't modify memory location which is holding original number
func reverseBigNum(num *big.Int) *big.Int {
	tmpN := big.NewInt(int64(0)).Set(num)
	rev := big.NewInt(int64(0))
	zero := big.NewInt(int64(0))
	ten := big.NewInt(int64(10))
	for tmpN.Cmp(zero) == 1 {
		rem := big.NewInt(int64(0))
		tmpN, rem = tmpN.DivMod(tmpN, ten, rem)
		rev = rev.Mul(rev, ten)
		rev = rev.Add(rev, rem)
	}
	return rev
}

// given a big integer, checks whether it is
// palindorme or not
func isPalindromeBig(num *big.Int) bool {
	return num.Cmp(reverseBigNum(num)) == 0
}

// given an integer , checks whether it's
// lychrel or not
//
// see definition of lychrel number [here](https://projecteuler.net/problem=55)
//
// communicates via given boolean channel
func isLychrelNumber(num int, channel chan bool) {
	check := true
	numBig := big.NewInt(int64(num))
	for i := 0; i < 50; i++ {
		numBig = numBig.Add(numBig, reverseBigNum(numBig))
		if isPalindromeBig(numBig) {
			check = false
			break
		}
	}
	channel <- check
}

// LychrelNumbers - Parallelly computes how many
// numbers below 10K are lychrel, leveraging power of
// multicore CPU using go-routines
func LychrelNumbers() int {
	count := 0                    // holds count of lychrel numbers
	channel := make(chan bool, 1) // communication channel
	var wg sync.WaitGroup
	wg.Add(1)   // adds worker
	go func() { // listener function
		c := 0
		for i := range channel {
			if i {
				count++
			}
			c++
			if c == 9999 {
				break
			}
		}
		close(channel)
		wg.Done()
	}()
	for i := 1; i < 10000; i++ {
		go isLychrelNumber(i, channel)
	}
	wg.Wait()
	return count
}
