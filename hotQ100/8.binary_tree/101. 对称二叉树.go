package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var slove func(*TreeNode, *TreeNode) bool
	slove = func(tn1, tn2 *TreeNode) bool {
		if tn1 != nil && tn2 == nil {
			return false
		} else if tn1 == nil && tn2 != nil {
			return false
		} else if tn1 == nil && tn2 == nil {
			return true
		} else if tn1.Val != tn2.Val {
			return false
		}
		in := slove(tn1.Left, tn2.Right)
		out := slove(tn1.Right, tn2.Left)
		return in && out
	}
	return slove(root.Left, root.Right)
}

func isSymmetric(root *TreeNode) bool {
	var solve func(node1, node2 *TreeNode) bool
	solve = func(node1, node2 *TreeNode) bool {
		if node1 == nil && node2 == nil {
			return true
		} else if node1 == nil && node2 != nil {
			return false
		} else if node1 != nil && node2 == nil {
			return false
		} else if node1.Val != node2.Val {
			return false
		}

		o := solve(node1.Left, node2.Right)
		i := solve(node1.Right, node2.Left)

		return o && i
	}
	return solve(root.Left, root.Right)
}
