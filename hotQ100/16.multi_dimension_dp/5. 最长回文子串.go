package multidimensiondp

/**
 * dp[i][j] 表示下标[i,j]是否是回文串
 * 当s[i] == s[j] 时，若[i+1, j-1]也是一个回文串，则[i,j]也是回文串
 *                   若 j-1 < i+1 即 j-i <= 1 时，[i,j]一定是回文串
 * 当s[i] != s[j]时，[i, j]不可能是回文串
 * i只会 <= j, 只需遍历矩阵左上半部分
 * 遍历顺序为从下到上，从左到右
 */
func longestPalindrome1(s string) string {
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

func longestPalindrome2(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	len := 0
	ans := ""
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] {
				if j-i <= 2 {
					dp[i][j] = true
				} else if dp[i+1][j-1] {
					dp[i][j] = true
				}
			}
			if dp[i][j] {
				if j-i+1 > len {
					ans = s[i : j+1]
					len = j - i + 1
				}
			}
		}
	}
	return ans
}

func longestPalindrome(s string) string {
	l_ans, r_ans := 0, 0
	n := len(s)

	for i := range n {
		l, r := i, i
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}
		if r-l > r_ans-l_ans {
			r_ans = r
			l_ans = l
		}
	}
	for i := range n - 1 {
		l, r := i, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}
		if r-l > r_ans-l_ans {
			r_ans = r
			l_ans = l
		}
	}
	return s[l_ans+1 : r_ans]
}
