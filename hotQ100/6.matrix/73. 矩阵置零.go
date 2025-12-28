package hot

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	firstRow, firstCol := false, false

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				if i == 0 {
					firstRow = true
				}
				if j == 0 {
					firstCol = true
				}
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if firstRow {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
	if firstCol {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
