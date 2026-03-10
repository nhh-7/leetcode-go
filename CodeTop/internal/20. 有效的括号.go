package internal

func isValid(s string) bool {
	stack := []rune{}
	for _, ch := range s {
		switch ch {
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, ch)
		}
	}
	return len(stack) == 0
}
