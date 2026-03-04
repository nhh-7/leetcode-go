package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth1(root.Left)
	right := maxDepth1(root.Right)

	return max(left, right) + 1
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	return max(left, right) + 1
}
