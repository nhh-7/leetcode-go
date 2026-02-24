package multidimensiondp

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	edge := 0
	for i := range m {
		dp[i][0] = int(matrix[i][0] - '0')
		if dp[i][0] == 1 {
			edge = 1
		}
	}
	for j := range n {
		dp[0][j] = int(matrix[0][j] - '0')
		if dp[0][j] == 1 {
			edge = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1], dp[i][j-1]) + 1
				edge = max(edge, dp[i][j])
			}
		}
	}
	return edge * edge
}
