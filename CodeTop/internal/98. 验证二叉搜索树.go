package internal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {

	var pre *TreeNode
	var dfs func(tn *TreeNode) bool
	dfs = func(tn *TreeNode) bool {
		if tn == nil {
			return true
		}
		if !dfs(tn.Left) {
			return false
		}
		if pre != nil && pre.Val >= tn.Val {
			return false
		}
		pre = tn
		if !dfs(tn.Right) {
			return false
		}
		return true
	}
	return dfs(root)
}
