package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
当前节点若有left，则记录right将left放到right的位置，原来的right放到该left子树的右下，若无则继续下一right
*/
func flatten(root *TreeNode) {
	cur := root
	for cur != nil {
		if cur.Left != nil {
			tmp := cur.Right
			cl := cur.Left
			for cl.Right != nil {
				cl = cl.Right
			}
			cl.Right = tmp
			cur.Right = cur.Left
			cur.Left = nil
		}
		cur = cur.Right
	}
}
