package projecteuler

func navigate(x int, y int, maxX int, maxY int) int {
	if x == maxX || y == maxY {
		return 1
	}
	return navigate(x+1, y, maxX, maxY) + navigate(x, y+1, maxX, maxY)
}

// LatticePath - Recursively computes possible number of paths for reaching
// bottom-right cell, while starting at top-left cell (0, 0)
func LatticePath(x int, y int) int {
	return navigate(0, 0, x, y)
}
