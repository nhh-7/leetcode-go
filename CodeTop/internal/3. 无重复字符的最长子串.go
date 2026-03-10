package internal

func lengthOfLongestSubstring(s string) int {
	ans := 0
	cnt := [128]int{}

	left := 0
	for right := 0; right < len(s); right++ {
		cnt[s[right]]++
		for cnt[s[right]] > 1 {
			cnt[s[left]]--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}
