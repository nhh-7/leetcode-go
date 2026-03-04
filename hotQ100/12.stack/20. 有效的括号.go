package stack

func isValid1(s string) bool {
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

func isValid(s string) bool {
	stk := []rune{}
	for _, ch := range s {
		switch ch {
		case '(':
			stk = append(stk, ')')
		case '{':
			stk = append(stk, '}')
		case '[':
			stk = append(stk, ']')
		default:
			if len(stk) > 0 && ch == stk[len(stk)-1] {
				stk = stk[:len(stk)-1]
			} else {
				return false
			}
		}
	}
	return len(stk) == 0
}
