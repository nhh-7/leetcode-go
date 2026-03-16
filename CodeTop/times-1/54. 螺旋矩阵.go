package times1

func spiralOrder(matrix [][]int) []int {
	DIRS := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	m, n := len(matrix), len(matrix[0])
	res := make([]int, m*n)

	k := 0
	i, j := 0, 0
	d := 0
	for k < m*n {
		res[k] = matrix[i][j]
		matrix[i][j] = 101
		ni, nj := i+DIRS[d][0], j+DIRS[d][1]
		if ni < 0 || ni >= m || nj < 0 || nj >= n || matrix[ni][nj] == 101 {
			d = (d + 1) % 4
		}
		i = i + DIRS[d][0]
		j = j + DIRS[d][1]
		k++
	}
	return res
}
