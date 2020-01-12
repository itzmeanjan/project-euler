package projecteuler

import "sync"

// GetXthTriangularNumber - Returns triangular number at given position `x`
func GetXthTriangularNumber(x int) int {
	return x * (x + 1) / 2
}

// GetFactorCount - Return number of factors a given number have
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

// TriangularNumber - Holds a triangular number, along with number of its factors
type TriangularNumber struct {
	num     int
	factorC int
}

// this function tries to find out maximum triangular number
// with in a given range ( i.e. denotes position of triangular numbers )
//
// we're also having two communication channels, which will help us
// in sending back computed value & checking whether to abort computation or not, respectively
//
// `maxFactorC` - denotes maximum count of factors, from all triangular numbers of this given range ( by position of triangular number )
// `channel` - used to send computed result back to listener
// `channelEnd` - used for monitoring whether this goroutine is asked to abort its computation immediately or not
func getHighlyDivisibleTriNumFromRange(init int, end int, maxFactorC int, channel chan TriangularNumber, channelEnd chan bool) {
	triNum := TriangularNumber{1, 1}
	for i := init; i <= end; i++ {
		select {
		case <-channelEnd:
			return
		default:
			tmpX := GetXthTriangularNumber(i)
			if tmp := GetFactorCount(tmpX); tmp > triNum.factorC {
				triNum.factorC = tmp
				triNum.num = tmpX
				if triNum.factorC >= maxFactorC {
					break
				}
			}
		}
	}
	// pushing value ( triangular number with maximum number of factors, from given range ) via channel
	channel <- triNum
	if triNum.factorC >= maxFactorC {
		close(channel) // if found desired value, we go for closing this channel, to let listener know, there's nothing more to read
	}
}

// HighlyDivisibleTriangularNumber - Finds out first triangular number, which has
// factors >=500
func HighlyDivisibleTriangularNumber() int {
	triNum := TriangularNumber{1, 1}
	channel := make(chan TriangularNumber, 2)
	channelEnd := make(chan bool)
	startAt := 1
	const workerC = 4
	// this anonymous function is designed to create requested number of
	// goroutines, by invoking a certain function, with varied arguments
	createWorkers := func(startAt *int, incrBy int, c int) {
		for i := 0; i < c; i++ {
			go getHighlyDivisibleTriNumFromRange(*startAt, *startAt+incrBy-1, 500, channel, channelEnd)
			*startAt += incrBy
		}
	}
	var wg sync.WaitGroup
	wg.Add(1)
	// anonymous function, keeps track of computed values, from completed goroutines,
	// if not reached expected value, can create more goroutines, for perfroming next stage
	// of computation
	go func() {
		defer wg.Done()
		respC := 0 // keeps track how many worker has completed, upto this point
		for v := range channel {
			if v.factorC > triNum.factorC {
				triNum = v
			}
			if triNum.factorC >= 500 {
				channelEnd <- true
				break
			}
			respC++
			if respC%workerC == 0 {
				createWorkers(&startAt, 500, workerC)
			}
		}
	}()
	// initially creating specified number of goroutines, later on, if needed, we may request more workers ( for next stage of computation )
	createWorkers(&startAt, 500, workerC)
	wg.Wait() // blocking wait, until worker denotes work is completed
	return triNum.num
}
