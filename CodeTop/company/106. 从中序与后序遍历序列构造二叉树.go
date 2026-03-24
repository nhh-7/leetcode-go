package main

import "slices"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	val := postorder[len(postorder)-1]
	leftSize := slices.Index(inorder, val)
	root := &TreeNode{Val: val}
	root.Left = buildTree(inorder[:leftSize], postorder[:leftSize])
	root.Right = buildTree(inorder[leftSize+1:], postorder[leftSize:len(postorder)-1])
	return root
}
