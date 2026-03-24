package times2

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	path := []string{}
	ans := []string{}
	n := len(s)

	var backtrack func(start int)
	backtrack = func(start int) {
		if len(path) == 4 {
			if start == n {
				ans = append(ans, strings.Join(path, "."))
			}
			return
		}

		for i := 1; i <= 3; i++ {
			if start+i > n {
				continue
			}
			sep := s[start : start+i]
			if isValid(sep) {
				path = append(path, sep)
				backtrack(start + i)
				path = path[:len(path)-1]
			}
		}
	}
	backtrack(0)
	return ans
}

func isValid(sep string) bool {
	if len(sep) > 1 && sep[0] == '0' {
		return false
	}
	v, _ := strconv.Atoi(sep)
	return v <= 255
}
