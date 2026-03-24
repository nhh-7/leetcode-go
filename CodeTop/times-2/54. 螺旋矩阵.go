package times2

func spiralOrder(matrix [][]int) []int {
	DIRS := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	d := 0
	m, n := len(matrix), len(matrix[0])
	i, j := 0, 0
	k := 0
	ans := make([]int, m*n)
	for k < m*n {
		ans[k] = matrix[i][j]
		matrix[i][j] = 101

		ni, nj := i+DIRS[d][0], j+DIRS[d][1]
		if ni < 0 || ni >= m || nj < 0 || nj >= n || matrix[ni][nj] > 100 {
			d = (d + 1) % 4
		}
		i = i + DIRS[d][0]
		j = j + DIRS[d][1]
		k++
	}
	return ans
}
