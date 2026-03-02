package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
	var dfs func(*TreeNode)
	sum := 0
	dfs = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		dfs(tn.Right)
		sum += tn.Val
		tn.Val = sum
		dfs(tn.Left)
	}
	dfs(root)
	return root
}
