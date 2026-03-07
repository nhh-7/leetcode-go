package hot

func isCovered(cnt_s, cnt_t []int) bool {
	for i := 'a'; i <= 'z'; i++ {
		if cnt_s[i] < cnt_t[i] {
			return false
		}
	}
	for i := 'A'; i <= 'Z'; i++ {
		if cnt_s[i] < cnt_t[i] {
			return false
		}
	}
	return true
}

func minWindow1(s string, t string) string {
	var cnt_s, cnt_t [128]int
	for _, c := range t {
		cnt_t[c]++
	}

	ansLeft, ansRight := -1, len(s)
	left := 0
	for right, c := range s {
		cnt_s[c]++
		for isCovered(cnt_s[:], cnt_t[:]) {
			if (ansRight - ansLeft) > (right - left) {
				ansRight = right
				ansLeft = left
			}
			cnt_s[s[left]]--
			left++
		}
	}
	if ansLeft < 0 {
		return ""
	}
	return s[ansLeft : ansRight+1]
}

func minWindow(s string, t string) string {
	cnt_s, cnt_t := [128]int{}, [128]int{}
	for _, ch := range t {
		cnt_t[ch]++
	}

	left := 0
	ans_l, ans_r := -1, len(s)
	for right, ch := range s {
		cnt_s[ch]++
		for isCovered(cnt_s[:], cnt_t[:]) {
			if right-left < ans_r-ans_l {
				ans_l = left
				ans_r = right
			}
			cnt_s[s[left]]--
			left++
		}
	}
	if ans_l < 0 {
		return ""
	}
	return s[ans_l : ans_r+1]
}
