package backtrack

func generateParenthesis1(n int) []string {
	path := []byte{}
	ans := []string{}

	var dfs func(int, int)
	dfs = func(left, right int) {
		if right == n {
			ans = append(ans, string(path))
		}
		if left < n {
			path = append(path, '(')
			dfs(left+1, right)
			path = path[:len(path)-1]
		}

		if right < left {
			path = append(path, ')')
			dfs(left, right+1)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	return ans
}

func generateParenthesis(n int) []string {
	ans := []string{}
	path := []byte{}

	var backtrack func(int, int)
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
