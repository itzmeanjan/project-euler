package projecteuler

// dynamic programming based solution - improved version over recursion based approach
func navigateI(x int, y int) int {
	// initializing a (x+1)*(y+1) 2D slice
	mem := make([][]int, x+1)
	for i := range mem {
		mem[i] = make([]int, y+1)
	}
	// intializing target location / index with possible path count 1
	mem[x][y] = 1
	// computing path count via indices ( to x, y index ) along bottom-most row
	for i := x - 1; i >= 0; i-- {
		mem[i][y] += mem[i+1][y]
	}
	// computing path count via indices ( to x, y index ) along right-most column
	for i := y - 1; i >= 0; i-- {
		mem[x][i] += mem[x][i+1]
	}
	// computing possible path count via all remaining indices, which are not on edge or border
	for i := x - 1; i >= 0; i-- {
		for j := y - 1; j >= 0; j-- {
			mem[i][j] = mem[i+1][j] + mem[i][j+1]
		}
	}
	return mem[0][0]
}

// a very computationally intensive recursive solution
func navigate(x int, y int, maxX int, maxY int) int {
	if x == maxX || y == maxY {
		return 1
	}
	return navigate(x+1, y, maxX, maxY) + navigate(x, y+1, maxX, maxY)
}

// LatticePath - Recursively computes possible number of paths for reaching
// bottom-right cell, while starting at top-left cell (0, 0)
func LatticePath(x int, y int) int {
	return navigateI(x, y)
	//return navigate(0, 0, x, y)
}
