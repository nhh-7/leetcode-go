package backtrack

func letterCombinations(digits string) []string {
	mapping := [...]string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	ans := []string{}
	path := []byte{}
	n := len(digits)
	var dfs func(int)
	dfs = func(index int) { // index 直接代表当前处理 digits[index]
		if index == n { // 终止条件：处理完了所有数字
			ans = append(ans, string(path))
			return
		}

		// 取出当前数字对应的字母串
		letters := mapping[digits[index]-'0']

		// 遍历当前数字能选的字母
		for j := 0; j < len(letters); j++ {
			path = append(path, letters[j])
			dfs(index + 1)            // 递归处理【下一个】数字
			path = path[:len(path)-1] // 回溯
		}
	}
	dfs(0)
	return ans
}
