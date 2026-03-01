package dp

func numSquares1(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i
	}

	for i := 1; i*i <= n; i++ {
		for j := i * i; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-i*i]+1)
		}
	}
	return dp[n]
}

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i // 初始化全部由1组成
	}
	dp[0] = 0
	for i := 1; i*i <= n; i++ {
		for j := i * i; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-i*i]+1)
		}
	}
	return dp[n]
}
