package projecteuler

import (
	"math"
	"sync"
)

type rightTriangleCountHolder struct {
	count int
	value int
}

// custom type holding side lengths of a right triangle
// where a < b < c, needs to be satisfied
type rightTriangle struct {
	a int
	b int
	c int
}

// checks whether given `rightTriangle`
// satisfies pythagorean theorem or not
// a^2 + b^2 == c^2, a < b < c
func (t rightTriangle) isValid() bool {
	square := func(b int) int {
		return int(math.Pow(float64(b), 2.0))
	}
	return square(t.a)+square(t.b) == square(t.c)
}

// calculates number of right triangles that can be
// formed using three side (a, b, c) combination, while
// satisfying a + b + c == p && a < b < c && a^2 + b^2 == c^2
func rightTriangleCount(p int, channel chan rightTriangleCountHolder) {
	buffer := make([]rightTriangle, 0)
	for a := 1; a < p; a++ {
		for b := a + 1; b < p; b++ {
			if a+b >= p {
				break
			}
			for c := b + 1; c < p; c++ {
				if sum := a + b + c; sum > p {
					break
				} else if sum < p {
					continue
				}
				t := rightTriangle{a, b, c}
				if t.isValid() {
					buffer = append(buffer, t)
				}
			}
		}
	}
	channel <- rightTriangleCountHolder{len(buffer), p}
}

// IntegerRightTriangles - Finds that value between 3 and 1000,
// for which maximum number of right triangles can be formed, using
// some triangle side combination, satisfying a + b + c == p, 3 <= p < 1001
func IntegerRightTriangles() int {
	max := rightTriangleCountHolder{0, 0}
	channel := make(chan rightTriangleCountHolder, 1)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		c := 0
		for i := range channel {
			if i.count > max.count {
				max = i
			}
			c++
			if c == 1000-3+1 {
				break
			}
		}
		close(channel)
		wg.Done()
	}()
	for p := 3; p < 1001; p++ {
		go rightTriangleCount(p, channel)
	}
	wg.Wait()
	return max.value
}
