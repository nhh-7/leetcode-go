package times2

func isCovered(cntS, cntT []int) bool {
	for i := 'a'; i <= 'z'; i++ {
		if cntS[i] < cntT[i] {
			return false
		}
	}
	for i := 'A'; i <= 'Z'; i++ {
		if cntS[i] < cntT[i] {
			return false
		}
	}
	return true
}

func minWindow(s string, t string) string {
	cntS, cntT := [128]int{}, [128]int{}

	for _, ch := range t {
		cntT[ch]++
	}

	left := 0
	start := -1
	minLen := len(s) + 1
	for right, ch := range s {
		cntS[ch]++
		for isCovered(cntS[:], cntT[:]) {
			if right-left+1 < minLen {
				minLen = right - left + 1
				start = left
			}
			cntS[s[left]]--
			left++
		}
	}
	if start == -1 {
		return ""
	}
	return s[start : start+minLen]
}
