package backtrack

func generateParenthesis(n int) []string {
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
