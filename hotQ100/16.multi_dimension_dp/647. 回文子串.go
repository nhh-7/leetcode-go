package multidimensiondp

// / 定义 `dp[i][j] `表示子串 s[i..j] 是否为回文，若 s[i] == s[j] 且满足“子串长度 <= 2”或“内部子串` dp[i+1][j-1] `为真”，则当前子串为回文，最终累加所有为 true 的状态。
func countSubstrings(s string) int {
	ans := 0
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] {
				if j-i <= 1 || dp[i+1][j-1] {
					dp[i][j] = true
					ans++
				}
			}
		}
	}
	return ans
}
