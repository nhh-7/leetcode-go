package internal

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	edge := 0
	for i := 0; i < m; i++ {
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			edge = 1
		}
	}
	for j := 0; j < n; j++ {
		if matrix[0][j] == '1' {
			dp[0][j] = 1
			edge = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
			edge = max(edge, dp[i][j])
		}
	}
	return edge * edge
}
