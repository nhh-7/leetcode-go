package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	m := len(nums) / 2
	node := &TreeNode{
		Val: nums[m],
	}
	node.Left = sortedArrayToBST(nums[:m])
	node.Right = sortedArrayToBST(nums[m+1:])
	return node
}
