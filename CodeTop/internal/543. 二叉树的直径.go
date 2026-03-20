package internal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	var dfs func(*TreeNode) int
	ans := 0
	dfs = func(tn *TreeNode) int {
		if tn == nil {
			return 0
		}
		left := dfs(tn.Left)
		right := dfs(tn.Right)
		ans = max(ans, left+right+1)
		return max(left, right) + 1
	}
	dfs(root)
	return ans - 1
}
