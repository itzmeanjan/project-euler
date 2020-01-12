package projecteuler

import (
	"fmt"
	"sync"
)

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
	channel <- triNum
	if triNum.factorC >= maxFactorC {
		close(channel)
	}
}

// HighlyDivisibleTriangularNumber - ...
func HighlyDivisibleTriangularNumber() int {
	triNum := TriangularNumber{1, 1}
	channel := make(chan TriangularNumber, 2)
	channelEnd := make(chan bool)
	startAt := 1
	workerC := 4
	createWorkers := func(startAt *int, incrBy int, c int) {
		for i := 0; i < c; i++ {
			go getHighlyDivisibleTriNumFromRange(*startAt, *startAt+incrBy-1, 500, channel, channelEnd)
			*startAt += incrBy
		}
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		respC := 1
		for v := range channel {
			if v.factorC > triNum.factorC {
				triNum = v
			}
			fmt.Println(v)
			if triNum.factorC >= 500 {
				channelEnd <- true
				wg.Done()
			}
			respC++
			if respC%workerC == 0 {
				createWorkers(&startAt, 250, workerC)
			}
		}
	}()
	createWorkers(&startAt, 500, workerC)
	wg.Wait()
	return triNum.num
}
