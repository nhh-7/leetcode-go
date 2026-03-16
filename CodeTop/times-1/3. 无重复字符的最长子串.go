package times1

func lengthOfLongestSubstring(s string) int {
	cnt_s := [128]int{}

	left := 0
	ans := 0
	for right, ch := range s {
		cnt_s[ch]++
		for cnt_s[ch] > 1 {
			cnt_s[s[left]]--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}
