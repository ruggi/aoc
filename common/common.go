package common

func GaussJordanElimination(matrix [][]int, numVars int) ([][]int, []int) {
	rows := len(matrix)
	cols := numVars + 1 // +1 for the result column

	// Create a copy to work with
	m := make([][]int, rows)
	for i := range matrix {
		m[i] = make([]int, cols)
		copy(m[i], matrix[i])
	}

	pivotRow := 0
	pivotCols := []int{} // Track which columns have pivots

	for col := 0; col < numVars && pivotRow < rows; col++ {
		// Find pivot (non-zero element)
		foundPivot := -1
		for row := pivotRow; row < rows; row++ {
			if m[row][col] != 0 {
				foundPivot = row
				break
			}
		}

		if foundPivot == -1 {
			continue // This column is a free variable
		}

		// Swap rows to bring pivot to current position
		m[pivotRow], m[foundPivot] = m[foundPivot], m[pivotRow]

		pivotCols = append(pivotCols, col)
		pivotValue := m[pivotRow][col]

		// Eliminate all other rows using exact arithmetic
		for row := 0; row < rows; row++ {
			if row == pivotRow {
				continue
			}
			if m[row][col] != 0 {
				// To eliminate m[row][col], we need:
				// m[row] = m[row] * pivotValue - m[row][col] * m[pivotRow]
				// This keeps everything as integers
				rowCoef := m[row][col]
				for c := 0; c < cols; c++ {
					m[row][c] = m[row][c]*pivotValue - rowCoef*m[pivotRow][c]
				}

				// Simplify row by dividing by GCD
				g := 0
				for c := 0; c < cols; c++ {
					if m[row][c] != 0 {
						g = greatestCommonDivisor(g, Abs(m[row][c]))
					}
				}
				if g > 1 {
					for c := 0; c < cols; c++ {
						m[row][c] /= g
					}
				}
			}
		}

		// Simplify pivot row by dividing by GCD
		g := 0
		for c := 0; c < cols; c++ {
			if m[pivotRow][c] != 0 {
				g = greatestCommonDivisor(g, Abs(m[pivotRow][c]))
			}
		}
		if g > 1 {
			for c := 0; c < cols; c++ {
				m[pivotRow][c] /= g
			}
		}

		pivotRow++
	}

	// Determine free variables
	freeVars := []int{}
	pivotColsMap := make(map[int]bool)
	for _, pc := range pivotCols {
		pivotColsMap[pc] = true
	}
	for col := range numVars {
		if !pivotColsMap[col] {
			freeVars = append(freeVars, col)
		}
	}

	return m, freeVars
}

func greatestCommonDivisor(a, b int) int {
	a = Abs(a)
	b = Abs(b)
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
