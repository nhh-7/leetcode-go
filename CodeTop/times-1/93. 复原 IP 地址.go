package times1

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	ans := []string{}
	n := len(s)
	path := []string{}

	var backtrack func(int)
	backtrack = func(start int) {
		// if len(path) == 4 {
		// 	if start == n {
		// 		ans = append(ans, strings.Join(path, "."))
		// 	}
		// 	return
		// }
		if start == n {
			if len(path) == 4 {
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

func isValid(s string) bool {
	if len(s) > 1 && s[0] == '0' {
		return false
	}
	n, _ := strconv.Atoi(s)
	return n <= 255
}
