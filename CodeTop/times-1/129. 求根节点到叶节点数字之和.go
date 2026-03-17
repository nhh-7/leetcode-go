package times1

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {

	var dfs func(tn *TreeNode, sum int) int
	dfs = func(tn *TreeNode, sum int) int {
		if tn == nil {
			return 0
		}
		sum = sum*10 + tn.Val
		if tn.Left == nil && tn.Right == nil {
			return sum
		}
		left := dfs(tn.Left, sum)
		right := dfs(tn.Right, sum)
		return left + right
	}
	return dfs(root, 0)
}
