package internal

import "math"

func isCovered(cnt_s, cnt_t []int) bool {
	for i := 'a'; i <= 'z'; i++ {
		if cnt_t[i] > cnt_s[i] {
			return false
		}
	}
	for i := 'A'; i <= 'Z'; i++ {
		if cnt_t[i] > cnt_s[i] {
			return false
		}
	}
	return true
}

func minWindow(s string, t string) string {
	cnt_s, cnt_t := make([]int, 128), make([]int, 128)
	for _, ch := range t {
		cnt_t[ch]++
	}

	start, ansLen := -1, math.MaxInt
	left := 0
	for right, ch := range s {
		cnt_s[ch]++
		for isCovered(cnt_s, cnt_t) {
			if ansLen > right-left+1 {
				ansLen = right - left + 1
				start = left
			}
			cnt_s[s[left]]--
			left++
		}
	}
	if start == -1 {
		return ""
	}
	return s[start : start+ansLen]
}
