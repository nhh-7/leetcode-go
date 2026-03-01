package dp

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	var dfs func(*TreeNode) []int
	dfs = func(tn *TreeNode) []int {
		if tn == nil {
			return []int{0, 0}
		}
		left := dfs(tn.Left)
		right := dfs(tn.Right)

		// 情况 1: 选中当前节点，则左右孩子都不能选
		selected := tn.Val + left[0] + right[0]

		// 情况 2: 不选当前节点，则左右孩子选不选都行，取各自的最大值
		notSelected := max(left[0], left[1]) + max(right[0], right[1])

		return []int{notSelected, selected}
	}
	ans := dfs(root)
	return max(ans[0], ans[1])
}
