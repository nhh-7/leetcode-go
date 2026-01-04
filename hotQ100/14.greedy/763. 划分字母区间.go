package greedy

// 核心，同一字母只能出现在同一片段中，从前往后扫，知道当前片段中所有字母没有在后面出现了，记录答案
func partitionLabels(s string) []int {
	m := map[rune]int{}
	for i, c := range s {
		m[c] = i
	}

	ans := []int{}
	lst := -1
	sum := 0 // 已经划分了多少字符数
	for i, c := range s {
		lst = max(lst, m[c]) // 记录这一段中字符最后出现的索引位置

		// 这一段划分完了
		if lst == i {
			ans = append(ans, i-sum+1)
			sum += (i - sum + 1)
		}
	}
	return ans
}
