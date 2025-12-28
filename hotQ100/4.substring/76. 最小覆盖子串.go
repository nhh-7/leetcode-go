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

func minWindow(s string, t string) string {
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
