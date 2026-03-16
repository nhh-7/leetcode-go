package internal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	var dfs func(tn *TreeNode) int
	flag := true
	dfs = func(tn *TreeNode) int {
		if tn == nil {
			return 0
		}
		left := dfs(tn.Left)
		right := dfs(tn.Right)
		if abs(left-right) > 1 {
			flag = false
		}
		return max(left, right) + 1
	}
	dfs(root)
	return flag
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
