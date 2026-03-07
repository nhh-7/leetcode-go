package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST1(root *TreeNode) bool {
	var pre *TreeNode
	var traversal func(*TreeNode) bool
	traversal = func(tn *TreeNode) bool {
		if tn == nil {
			return true
		}
		if !traversal(tn.Left) {
			return false
		}
		if pre != nil && pre.Val >= tn.Val {
			return false
		}
		pre = tn
		if !traversal(tn.Right) {
			return false
		}
		return true
	}
	return traversal(root)
}

func isValidBST(root *TreeNode) bool {
	var pre *TreeNode = nil

	var dfs func(*TreeNode) bool
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
