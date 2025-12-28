package backtrack

func letterCombinations(digits string) []string {
	mapping := [...]string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	ans := []string{}
	path := []byte{}
	n := len(digits)
	if n == 0 {
		return ans
	}
	var dfs func(int)
	dfs = func(startIdx int) {
		if len(path) >= n {
			ans = append(ans, string(path))
			return
		}
		for i := startIdx; i < n; i++ {
			for j := 0; j < len(mapping[digits[i]-'0']); j++ {
				path = append(path, mapping[digits[i]-'0'][j])
				dfs(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return ans
}
