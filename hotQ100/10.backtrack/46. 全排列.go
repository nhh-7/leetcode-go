package backtrack

func permute(nums []int) [][]int {
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
