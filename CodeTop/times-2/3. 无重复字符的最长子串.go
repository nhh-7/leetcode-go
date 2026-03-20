package times2

func lengthOfLongestSubstring(s string) int {
	cnt := make([]int, 128)

	left := 0
	ans := 0
	for right, ch := range s {
		cnt[ch]++
		for cnt[ch] > 1 {
			cnt[s[left]]--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}
