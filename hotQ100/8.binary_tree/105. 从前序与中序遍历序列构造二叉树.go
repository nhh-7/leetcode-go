package binarytree

import "slices"

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

	leftSize := slices.Index(inorder, preorder[0])

	node := &TreeNode{Val: preorder[0]}
	node.Left = buildTree(preorder[1:1+leftSize], inorder[:leftSize])
	node.Right = buildTree(preorder[leftSize+1:], inorder[1+leftSize:])
	return node
}
