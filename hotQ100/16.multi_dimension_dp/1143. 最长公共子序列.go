package multidimensiondp

/*
dp[i][j]表示 下标到i-1的text1 和 下标到j-1的text2 的最长公共子序列
*/
func longestCommonSubsequence(text1 string, text2 string) int {
	n, m := len(text1), len(text2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	dp[0][0] = 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp[n][m]
}
