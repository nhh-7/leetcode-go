package internal

func generateParenthesis(n int) []string {
	ans := []string{}
	path := []byte{}

	var backtrack func(left, right int)
	backtrack = func(left, right int) {
		if left == n && right == n {
			ans = append(ans, string(path))
			return
		}

		if left < n {
			path = append(path, '(')
			backtrack(left+1, right)
			path = path[:len(path)-1]
		}

		if right < left {
			path = append(path, ')')
			backtrack(left, right+1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0, 0)
	return ans
}
