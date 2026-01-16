package multidimensiondp

/**
 * dp[i][j] 表示下标[i,j]是否是回文串
 * 当s[i] == s[j] 时，若[i+1, j-1]也是一个回文串，则[i,j]也是回文串
 *                   若 j-1 < i+1 即 j-i <= 1 时，[i,j]一定是回文串
 * 当s[i] != s[j]时，[i, j]不可能是回文串
 * i只会 <= j, 只需遍历矩阵左上半部分
 * 遍历顺序为从下到上，从左到右
 */
func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}

	len, start := -1, 0
	for i := n - 2; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					dp[i][j] = true
					if j-i > len {
						len = j - i
						start = i
					}
				} else if dp[i+1][j-1] {
					dp[i][j] = true
					if j-i > len {
						len = j - i
						start = i
					}
				}
			}
		}
	}
	if len == -1 {
		return string(s[0])
	}
	return s[start : start+len+1]
}
