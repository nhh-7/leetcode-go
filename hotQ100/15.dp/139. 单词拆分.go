package dp

/**
 * 排列问题，而不是组合
 * dp[j] 表示以j为长度的字符串是否可以字典的单词组成，true或false
 * 外层容量，内层物品
 * 逻辑：leetcode, 遍历到当前容量为3(j)时，挨个获取以s[2]结尾的子串，
 * 比如 lee, ee, e, 是否存在Dict中，如果存在，那么看dp[0]，dp[1]或dp[2]是否为true，为true则可以拼接上Dict中的串，从而dp[3]为true
 */
func wordBreak(s string, wordDict []string) bool {
	set := map[string]bool{}
	for _, word := range wordDict {
		set[word] = true
	}
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for j := 1; j <= n; j++ {
		for i := 0; i < j; i++ {
			sub := s[i:j]
			if set[sub] && dp[i] {
				dp[j] = true
			}
		}
	}
	return dp[n]
}
