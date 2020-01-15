package projecteuler

import (
	"fmt"
)

// Row - ...
type Row struct {
	elems     []int
	isUpdated []bool
	updated   []int
}

func (r Row) String() string {
	str := ""
	for _, i := range r.elems {
		str += fmt.Sprintf(" %d", i)
	}
	str += "\n"
	return str
}

// Triangle - ...
type Triangle struct {
	rows []Row
}

func (t Triangle) String() string {
	str := ""
	for _, i := range t.rows {
		str += i.String()
	}
	return str
}

func (t *Triangle) insertRows(elems [][]int) Triangle {
	copyarr := func(arr []int) []int {
		arrN := make([]int, len(arr))
		for i, j := range arr {
			arrN[i] = j
		}
		return arrN
	}
	for i, j := range elems {
		t.rows[i] = Row{j, make([]bool, i+1), copyarr(j)}
	}
	return *t
}

func buildTriangle() Triangle {
	// [][]int{{3}, {7, 4}, {2, 4, 6}, {8, 5, 9, 3}}
	data := [][]int{{75},
		{95, 64},
		{17, 47, 82},
		{18, 35, 87, 10},
		{20, 4, 82, 47, 65},
		{19, 1, 23, 75, 3, 34},
		{88, 2, 77, 73, 7, 63, 67},
		{99, 65, 4, 28, 6, 16, 70, 92},
		{41, 41, 26, 56, 83, 40, 80, 70, 33},
		{41, 48, 72, 33, 47, 32, 37, 16, 94, 29},
		{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14},
		{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57},
		{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48},
		{63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31},
		{4, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 4, 23}}
	tri := Triangle{make([]Row, len(data))}
	return tri.insertRows(data)
}

type neighbourStat struct {
	isUpdated bool
	row       int
	col       int
}

func (t Triangle) neighbours(i int, j int) [2]neighbourStat {
	return [2]neighbourStat{
		neighbourStat{t.rows[i+1].isUpdated[j], i + 1, j},
		neighbourStat{t.rows[i+1].isUpdated[j+1], i + 1, j + 1}}
}

// MaxPathSum - ...
func MaxPathSum() int {
	max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	maxPathSum := func(t Triangle) int {
		max := 0
		for _, i := range t.rows[len(t.rows)-1].updated {
			if max < i {
				max = i
			}
		}
		return max
	}
	tri := buildTriangle()
	for i := 0; i < len(tri.rows)-1; i++ {
		for j, k := range tri.rows[i].updated {
			for _, m := range tri.neighbours(i, j) {
				if m.isUpdated {
					tri.rows[m.row].updated[m.col] = max(tri.rows[m.row].updated[m.col], tri.rows[m.row].elems[m.col]+k)
				} else {
					tri.rows[m.row].updated[m.col] = tri.rows[m.row].elems[m.col] + k
					tri.rows[m.row].isUpdated[m.col] = true
				}
			}
		}
	}
	return maxPathSum(tri)
}
