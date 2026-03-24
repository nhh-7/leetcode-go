package times2

func longestValidParentheses(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = 0
	ans := 0
	for i := 1; i < n; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				dp[i] = 2
				if i-2 >= 0 && dp[i-2] > 0 {
					dp[i] += dp[i-2]
				}
			} else {
				if dp[i-1] > 0 && i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
					dp[i] = dp[i-1] + 2
					if i-dp[i-1]-2 >= 0 && dp[i-dp[i-1]-2] > 0 {
						dp[i] += dp[i-dp[i-1]-2]
					}
				}
			}
		} else {
			dp[i] = 0
		}
		ans = max(dp[i], ans)
	}
	return ans
}
