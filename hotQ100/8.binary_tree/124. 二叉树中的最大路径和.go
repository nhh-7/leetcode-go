package binarytree

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	var traversal func(*TreeNode) int
	var ans int = math.MinInt
	// traversal 返回tn到其子树下某节点的最大路径和
	traversal = func(tn *TreeNode) int {
		if tn == nil {
			return 0
		}
		left := traversal(tn.Left)
		right := traversal(tn.Right)
		ans = max(ans, left+tn.Val, right+tn.Val, left+right+tn.Val, tn.Val)
		return max(left, right, 0) + tn.Val
	}
	traversal(root)
	return ans
}
