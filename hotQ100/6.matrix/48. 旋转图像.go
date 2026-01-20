package hot

import "slices"

func rotate1(matrix [][]int) {
	n := len(matrix)

	rn := n / 2
	if n%2 == 1 {
		rn++
	}
	// [i][j] -> [j][n-i-1]
	// [n-j-1][i] -> [i][j]
	for i := 0; i < n/2; i++ {
		for j := 0; j < rn; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[n-j-1][i]
			matrix[n-j-1][i] = matrix[n-i-1][n-j-1]
			matrix[n-i-1][n-j-1] = matrix[j][n-i-1]
			matrix[j][n-i-1] = tmp
		}
	}
}

// 先将矩阵**沿主对角线（左上到右下）转置（交换）**，再将每一行**水平翻转**。
func rotate(matrix [][]int) {
	n := len(matrix)
	for i := range matrix {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for _, row := range matrix {
		slices.Reverse(row)
	}
}
