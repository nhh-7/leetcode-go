package hot

import "math"

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])

	ans := make([]int, m*n)

	DIRS := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	d := 0
	x, y := 0, 0

	for k := 0; k < m*n; k++ {

		ans[k] = matrix[x][y]
		matrix[x][y] = math.MaxInt

		tx := x + DIRS[d][0]
		ty := y + DIRS[d][1]

		if tx < 0 || tx >= m || ty < 0 || ty >= n || matrix[tx][ty] == math.MaxInt {
			d = (d + 1) % 4
		}

		x += DIRS[d][0]
		y += DIRS[d][1]
	}
	return ans
}
