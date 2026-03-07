package hot

func findAnagrams1(s string, p string) []int {
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

func findAnagrams(s string, p string) []int {
	cnt_s, cnt_p := [26]int{}, [26]int{}
	for _, ch := range p {
		cnt_p[ch-'a']++
	}
	ans := []int{}
	left := 0
	for right, ch := range s {
		cnt_s[ch-'a']++
		if right-left+1 < len(p) {
			continue
		}
		// 数组类型可以直接比较，切片不能直接比
		if cnt_s == cnt_p {
			ans = append(ans, left)
		}
		cnt_s[s[left]-'a']--
		left++
	}
	return ans
}
