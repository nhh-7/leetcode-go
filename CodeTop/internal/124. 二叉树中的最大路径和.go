package internal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	ans := -1001

	var traversal func(*TreeNode) int
	traversal = func(tn *TreeNode) int {
		if tn == nil {
			return 0
		}
		left := max(traversal(tn.Left), 0)
		right := max(traversal(tn.Right), 0)

		ans = max(ans, left+right+tn.Val)
		return max(left, right) + tn.Val
	}

	traversal(root)
	return ans
}
