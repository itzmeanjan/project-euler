package projecteuler

import "sync"

// checks whether a given number is triangular or not
func isTriangular(num int) bool {
	check := false
	for i := 1; ; i++ {
		if tmp := i * (i + 1) / 2; tmp > num {
			break
		} else if tmp == num {
			check = true
			break
		}
	}
	return check
}

// given a word formed of english alphabet,
// we'll calculate word value i.e. replacing each letter
// with its corresponding position value in alphabet ( starting from A -> 1 )
// and finally summing them up, of that word
func getWordValue(word string) int {
	val := 0
	for _, v := range word {
		val += (int(v) - 64)
	}
	return val
}

// we'll check whether word value of a given word is triangular or not,
// and send result using communication channel
func isTriangularWord(word string, channel chan bool) {
	channel <- isTriangular(getWordValue(word))
}

// CodedTriangleNumber - Finds how many words are triangular
// in a given word file, holding nearly 2K words
func CodedTriangleNumber() int {
	buffer := readFile("p042_words.txt")
	channel := make(chan bool, 1)
	count := 0
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		c := 0
		for i := range channel {
			c++
			if i {
				count++
			}
			if c == len(buffer) {
				break
			}
		}
		close(channel)
		wg.Done()
	}()
	for _, v := range buffer {
		go isTriangularWord(v, channel)
	}
	wg.Wait()
	return count
}
