package backtrack

func isPalidrome(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func partition(s string) [][]string {
	ans := [][]string{}
	path := []string{}
	var dfs func(int)
	n := len(s)
	dfs = func(startIdx int) {
		if startIdx == n {
			tmp := make([]string, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
		}

		for i := startIdx; i < n; i++ {
			if isPalidrome(s, startIdx, i) {
				path = append(path, s[startIdx:i+1])
				dfs(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return ans
}
