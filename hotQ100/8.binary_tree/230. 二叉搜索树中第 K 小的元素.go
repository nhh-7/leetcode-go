package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	var ans int
	var traversal func(*TreeNode) bool
	traversal = func(tn *TreeNode) bool {
		if tn == nil {
			return false
		}
		if traversal(tn.Left) {
			return true
		}
		k--
		if k == 0 {
			ans = tn.Val
			return true
		}
		if traversal(tn.Right) {
			return true
		}
		return false
	}
	traversal(root)
	return ans
}
