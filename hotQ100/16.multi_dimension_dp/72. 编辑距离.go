package multidimensiondp

/**
 * dp[i][j] 表示将[0, i-1]的word1 转换成 [0, j-1]的word2使用的最少操作数
 * word1[i-1] == word2[j-1]时 dp[i][j] = dp[i-1][j-1]
 * word1[i-1] != word2[j-1]时，有三种方法使其相等，
 *                            1.删除word1[i-1]比较[0,i-2]的word1和[0,j-1]的word2  dp[i][j] = dp[i-1][j] + 1
 *                            2.删除word2[j-1]比较[0,i-1]的word1和[0,j-2]的word2  dp[i][j] = dp[i][j-1] + 1
 *                            3.替换使得word1[i-1] == word2[j-1] dp[i][j] = dp[i-1][j-1] + 1
 */
func minDistance(word1 string, word2 string) int {
	n, m := len(word1), len(word2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for i := 0; i <= n; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= m; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[n][m]
}
