package internal

func permute(nums []int) [][]int {
	ans := [][]int{}
	path := []int{}
	n := len(nums)
	var backtrack func(map[int]struct{})
	backtrack = func(m map[int]struct{}) {
		if len(path) == n {
			c_path := make([]int, len(path))
			copy(c_path, path)
			ans = append(ans, c_path)
		}

		for _, v := range nums {
			if _, ok := m[v]; ok {
				continue
			}
			m[v] = struct{}{}
			path = append(path, v)
			backtrack(m)
			path = path[:len(path)-1]
			delete(m, v)
		}
	}
	backtrack(make(map[int]struct{}))
	return ans
}
