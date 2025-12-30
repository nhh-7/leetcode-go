package stack

func isValid(s string) bool {
	stk := make([]rune, 0)
	for _, b := range s {
		if b == '{' {
			stk = append(stk, '}')
		} else if b == '(' {
			stk = append(stk, ')')
		} else if b == '[' {
			stk = append(stk, ']')
		} else {
			if len(stk) > 0 && stk[len(stk)-1] == b {
				stk = stk[:len(stk)-1]
			} else {
				return false
			}
		}
	}
	return len(stk) == 0
}
