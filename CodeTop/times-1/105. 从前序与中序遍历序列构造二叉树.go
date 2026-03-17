package times1

import (
	"slices"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	val := preorder[0]
	leftSize := slices.Index(inorder, val)
	root := &TreeNode{Val: val}
	root.Left = buildTree(preorder[1:leftSize+1], inorder[:leftSize])
	root.Right = buildTree(preorder[leftSize+1:], inorder[leftSize+1:])
	return root
}
