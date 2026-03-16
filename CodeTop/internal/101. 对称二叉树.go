package internal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	var solve func(node1, node2 *TreeNode) bool
	solve = func(node1, node2 *TreeNode) bool {
		if node1 == nil && node2 != nil {
			return false
		} else if node1 != nil && node2 == nil {
			return false
		} else if node1 == nil && node2 == nil {
			return true
		} else if node1.Val != node2.Val {
			return false
		}
		in := solve(node1.Right, node2.Left)
		out := solve(node1.Left, node2.Right)
		return in && out
	}
	return solve(root.Left, root.Right)
}
