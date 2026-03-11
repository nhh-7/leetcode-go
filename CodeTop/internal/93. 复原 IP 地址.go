package internal

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	ans := []string{}
	path := []string{}

	var backtrack func(int)
	backtrack = func(start int) {
		if len(path) == 4 {
			if start == len(s) {
				ans = append(ans, strings.Join(path, "."))
			}
			return
		}

		for i := 1; i <= 3; i++ {
			if start+i > len(s) {
				break
			}
			segment := s[start : start+i]
			if valid(segment) {
				path = append(path, segment)
				backtrack(start + i)
				path = path[:len(path)-1]
			}
		}
	}
	backtrack(0)
	return ans
}

func valid(seg string) bool {
	if len(seg) > 1 && seg[0] == '0' {
		return false
	}
	num, _ := strconv.Atoi(seg)
	return num <= 255
}
