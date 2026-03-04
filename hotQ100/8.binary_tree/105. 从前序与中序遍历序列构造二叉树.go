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
func buildTree1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	leftSize := slices.Index(inorder, preorder[0])

	node := &TreeNode{Val: preorder[0]}
	node.Left = buildTree1(preorder[1:1+leftSize], inorder[:leftSize])
	node.Right = buildTree1(preorder[leftSize+1:], inorder[1+leftSize:])
	return node
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	val := preorder[0]
	root := &TreeNode{Val: val}
	idx := slices.Index(inorder, val)
	root.Left = buildTree(preorder[1:idx+1], inorder[:idx])
	root.Right = buildTree(preorder[idx+1:], inorder[idx+1:])
	return root
}
