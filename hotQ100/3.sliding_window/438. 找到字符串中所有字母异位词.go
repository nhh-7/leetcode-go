package hot

func findAnagrams(s string, p string) []int {
	ans := []int{}

	var cnt_p, cnt_s [26]int
	for _, c := range p {
		cnt_p[c-'a']++
	}

	left := 0
	for right, c := range s {
		cnt_s[c-'a']++
		if right-left+1 < len(p) {
			continue
		}
		if cnt_s == cnt_p {
			ans = append(ans, left)
		}
		cnt_s[s[left]-'a']--
		left++
	}
	return ans
}
