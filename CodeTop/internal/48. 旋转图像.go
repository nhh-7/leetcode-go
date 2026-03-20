package internal

import "slices"

func rotate(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])

	for i := 0; i < m; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for _, row := range matrix {
		slices.Reverse(row)
	}
}
