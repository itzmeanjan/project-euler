package projecteuler

// calculates sum of all elements present on both diagonals
// of a sqaure matrix ( 2D array/ slice )
func findSumOfItemsOnDiagonal(matrix *[][]int) int {
	mainDiagSum := 0
	for i, j := 0, 0; i < len(*matrix) && j < len((*matrix)[i]); {
		mainDiagSum += (*matrix)[i][j]
		i++
		j++
	}
	secDiagSum := 0
	for i, j := 0, len(*matrix)-1; i < len(*matrix) && j >= 0; {
		secDiagSum += (*matrix)[i][j]
		i++
		j--
	}
	return mainDiagSum + secDiagSum - 1
}

// given a sub-matrix of certain size ( i.e. width x width sub-matrix )
// we'll calculate all values in spiral fashion
//
// starting point of next iteration, with increased width ( increased by 2 )
// to be set from this function, cause we're passing value by reference
//
// even value to be put in next place to updated, which is also passed as reference
//
// `matrix` is updatable, cause we've passed it using reference
func fillInSpiralFields(width int, startI *int, startJ *int, startV *int, matrix *[][]int) {
	target := *startI + width - 1
	// moving towards down
	for *startI < target {
		(*matrix)[*startI][*startJ] = *startV
		*startV++
		*startI++
	}
	*startI--
	*startJ--
	target = *startJ - width + 1
	// moving towards left
	for *startJ > target {
		(*matrix)[*startI][*startJ] = *startV
		*startV++
		*startJ--
	}
	*startJ++
	*startI--
	target = *startI - width + 1
	// moving towards up
	for *startI > target {
		(*matrix)[*startI][*startJ] = *startV
		*startV++
		*startI--
	}
	*startI++
	*startJ++
	target = *startJ + width - 1
	// moving towards right
	for *startJ < target {
		(*matrix)[*startI][*startJ] = *startV
		*startV++
		*startJ++
	}
	// so, it's pretty well understandable that we're moving in clockwise fashion
}

// NumberSpiralDiagonals - Finds sum of numbers present on both diagonals of a 1001x1001
// matrix, which is formed using specified method, described in question
func NumberSpiralDiagonals(size int) int {
	matrix := make([][]int, size)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, size)
	} // initializes 2D array of given size, with initial 0 values in all fields
	startI := size / 2
	startJ := size / 2
	startV := 1
	matrix[startI][startJ] = startV // putting 1 at center i.e. cross-section of two diagonals
	startJ++
	startV++
	for i := 3; i <= size; i += 2 {
		fillInSpiralFields(i, &startI, &startJ, &startV, &matrix) // filling fields in spirally, in clockwise fashion
	}
	return findSumOfItemsOnDiagonal(&matrix)
}
