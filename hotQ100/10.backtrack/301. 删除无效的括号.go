package backtrack

func removeInvalidParentheses(s string) []string {
	lremove, rremove := 0, 0
	for _, ch := range s {
		if ch == '(' {
			lremove++
		} else if ch == ')' {
			if lremove > 0 {
				lremove--
			} else {
				rremove++
			}
		}
	}
	var res []string
	backtrack(s, 0, lremove, rremove, &res)
	return res
}

func valid(s string) bool {
	cnt := 0
	for _, ch := range s {
		if ch == '(' {
			cnt++
		} else if ch == ')' {
			cnt--
			if cnt < 0 {
				return false
			}
		}
	}
	return cnt == 0
}

func backtrack(s string, start, lremove, rremove int, res *[]string) {
	if lremove == 0 && rremove == 0 {
		if valid(s) {
			*res = append(*res, s)
		}
		return
	}

	for i := start; i < len(s); i++ {
		if i > start && s[i] == s[i-1] {
			continue
		}
		if lremove > 0 && s[i] == '(' {
			backtrack(s[:i]+s[i+1:], i, lremove-1, rremove, res)
		}
		if rremove > 0 && s[i] == ')' {
			backtrack(s[:i]+s[i+1:], i, lremove, rremove-1, res)
		}
	}
}
