package hot

func lengthOfLongestSubstring(s string) int {
	ans := 0
	n := len(s)
	cnt := [128]int{}

	left := 0
	for right := range n {
		cnt[s[right]]++

		for cnt[s[right]] > 1 {
			cnt[s[left]]--
			left++
		}

		ans = max(ans, right-left+1)
	}
	return ans
}
