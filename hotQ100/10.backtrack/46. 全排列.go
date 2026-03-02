package backtrack

func permute1(nums []int) [][]int {
	ans := [][]int{}
	n := len(nums)
	path := make([]int, 0)
	vis := make([]bool, n)

	var dfs func()
	dfs = func() {
		if n == len(path) {
			// 进行深拷贝，避免最后结果都是一样的
			tmp := make([]int, n)
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}
		for j := range n {
			if !vis[j] {
				path = append(path, nums[j])
				vis[j] = true
				dfs()
				vis[j] = false
				path = path[:len(path)-1]
			}
		}
	}
	dfs()
	return ans
}

func permute(nums []int) [][]int {
	ans := [][]int{}
	path := []int{}
	vis := make(map[int]bool)

	var backtrack func(map[int]bool)
	backtrack = func(vis map[int]bool) {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if !vis[i] {
				path = append(path, nums[i])
				vis[i] = true
				backtrack(vis)
				vis[i] = false
				path = path[:len(path)-1]
			}
		}
	}
	backtrack(vis)
	return ans
}
