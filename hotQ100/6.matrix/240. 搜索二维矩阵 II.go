package hot

func searchMatrix1(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])

	sm, sn := 0, n-1
	for sm < m && sn >= 0 {
		if matrix[sm][sn] > target {
			sn--
		} else if matrix[sm][sn] < target {
			sm++
		} else {
			return true
		}
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	row, col := 0, n-1
	for row < m && col >= 0 {
		if target < matrix[row][col] {
			col--
		} else if target > matrix[row][col] {
			row++
		} else {
			return true
		}
	}
	return false
}
