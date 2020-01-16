package projecteuler

import "fmt"

// holds original data present in a certain row
// along with their updated copy
type row struct {
	elems   []int
	updated []int
}

// string representation of a certain `row` in `Triangle`
func (r row) String() string {
	str := ""
	for _, i := range r.elems {
		str += fmt.Sprintf(" %d", i)
	}
	str += "\n"
	return str
}

// Triangle - holds whole Triangle, which is nothing but
// a slice of `row`-s
type Triangle struct {
	rows []row
}

// string representation of `Triangle`
func (t Triangle) String() string {
	str := ""
	for _, i := range t.rows {
		str += i.String()
	}
	return str
}

// inserts row(s) into Triangle i.e. puts data into Triangle, upon
// which we'll apply our algorithm for finding maximum cost path
func (t *Triangle) insertRows(elems [][]int) Triangle {
	copyarr := func(arr []int) []int {
		arrN := make([]int, len(arr))
		for i, j := range arr {
			arrN[i] = j
		}
		return arrN
	}
	for i, j := range elems {
		t.rows[i] = row{j, copyarr(j)}
	}
	return *t
}

// BuildTriangleI - builds Triangle from input data, and returns a Triangle
// upon which we'll apply our algorithm, for finding max cost path from top to bottom level
func BuildTriangleI() Triangle {
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
	tri := Triangle{make([]row, len(data))}
	return tri.insertRows(data)
}

// identifies neighboring positions in Triangle
// with `row` and `col` index, w.r.t. a certain position in Triangle
//
// `isUpdated` flag denotes, whether this field was updated in previous iteration or not
type neighbourStat struct {
	isUpdated bool
	row       int
	col       int
}

// returns two neighbors w.r.t. given indices of Triangle,
// present in next row of Triangle
func (t Triangle) neighbours(i int, j int) [2]neighbourStat {
	isUpdated := func(val1 int, val2 int) bool {
		if val1 != val2 {
			return true
		}
		return false
	}
	return [2]neighbourStat{
		neighbourStat{isUpdated(t.rows[i+1].elems[j], t.rows[i+1].updated[j]), i + 1, j},
		neighbourStat{isUpdated(t.rows[i+1].elems[j+1], t.rows[i+1].updated[j+1]), i + 1, j + 1}}
}

// MaxPathSum - Finds maximum cost path, when we start traversing from
// top of triangle and move towards bottom level
func MaxPathSum(tri Triangle) int {
	// computes maximum between two given numbers
	max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// finds maximum cost path in triangle from top to bottom level
	maxPathSum := func(t Triangle) int {
		max := 0
		for _, i := range t.rows[len(t.rows)-1].updated {
			if max < i {
				max = i
			}
		}
		return max
	}
	// updates cost of triangle, if we start moving from top to
	// bottom level. In each index, we can make two decisions, i.e.
	// which way to go, because from every position we'll eventually get
	// two neighbors ( in next row of triangle ), finally resulting into 2^14 = 16384
	// number of possible routes which can be followed from top to bottom
	//
	// and our job is to find out maximum cost path
	updateCost := func(tri Triangle) Triangle {
		for i := 0; i < len(tri.rows)-1; i++ {
			for j, k := range tri.rows[i].updated {
				for _, m := range tri.neighbours(i, j) {
					if m.isUpdated {
						tri.rows[m.row].updated[m.col] = max(tri.rows[m.row].updated[m.col], tri.rows[m.row].elems[m.col]+k)
					} else {
						tri.rows[m.row].updated[m.col] = tri.rows[m.row].elems[m.col] + k
					}
				}
			}
		}
		return tri
	}
	return maxPathSum(updateCost(tri))
}
