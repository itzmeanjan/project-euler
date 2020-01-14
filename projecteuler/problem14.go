package projecteuler

import "sync"

// holds property of collatz chain
// i.e. initial number of chain & length of chain
// produced ( numbers generated before we reach 1,
// because it's thought that all collatz seqs end at 1 )
type chainProp struct {
	init   int
	length int
}

// generates next term of collatz sequence, given current term
func generateNextTermOfCollatz(n int) int {
	if n%2 == 0 {
		return n / 2
	}
	return 3*n + 1
}

// generates collatz seqeuence for a given starting number
// and keeps track length of chain upto this point
//
// communicates with listener go-routine via channel
func generateCollatzSeq(cp chainProp, channel chan chainProp) {
	for n := cp.init; n != 1; {
		n = generateNextTermOfCollatz(n)
		cp.length++
	}
	channel <- cp
}

// LongestCollatzSeq - finds longest collatz sequence for a certain initial number
// i.e. for which initial number under 10^6, this sequence is of length max.
//
// deploying 10^6 lightweight go-routines, for computing desired value at
// lightning fast speed
func LongestCollatzSeq() int {
	maxChain := chainProp{1, 1}
	var wg sync.WaitGroup
	channel := make(chan chainProp, 2)
	wg.Add(1)
	go func() {
		c := 0
		for v := range channel {
			if v.length > maxChain.length {
				maxChain = v
			}
			c++
			if c == 999999 {
				close(channel)
				wg.Done()
			}
		}
	}()
	for cur := 1; cur < 1000000; cur++ {
		go generateCollatzSeq(chainProp{cur, 1}, channel)
	}
	wg.Wait()
	return maxChain.init
}
