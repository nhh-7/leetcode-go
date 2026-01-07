package dp

// dp[i] 表示到下标i有多少的最长有效括号子串长度
func longestValidParentheses(s string) int {
	n := len(s)
	dp := make([]int, n)
	if n == 0 {
		return 0
	}

	dp[0] = 0
	ans := 0
	for i := 1; i < n; i++ {
		if s[i] == '(' {
			dp[i] = 0
		} else {
			if s[i-1] == '(' {
				dp[i] = 2     // 匹配一对
				if i-2 >= 0 { // 类加上之前的有效括号
					dp[i] += dp[i-2]
				}
			} else {
				// s[i-1] == ')'情况下，如果s[i-1]都没有和前面的匹配上，那么s[i]更不可能匹配
				// 只有dp[i-1] > 0，且与s[i-1]匹配的'('的上一个字符（位置为 i - 1 - dp[i - 1]）同样是'(',s[i]就能和其匹配
				if dp[i-1] > 0 && i-1-dp[i-1] >= 0 && s[i-1-dp[i-1]] == '(' {
					dp[i] = 2 + dp[i-1]
					if i-2-dp[i-1] >= 0 { // 累加上一段
						dp[i] += dp[i-2-dp[i-1]]
					}
				}
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}
