package internal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	path := []int{}
	ans := [][]int{}

	var dfs func(tn *TreeNode, target int)
	dfs = func(tn *TreeNode, target int) {
		if tn == nil {
			return
		}
		path = append(path, tn.Val)
		if tn.Left == nil && tn.Right == nil {
			if tn.Val == target {
				tmp := make([]int, len(path))
				copy(tmp, path)
				ans = append(ans, tmp)
			}
		}
		dfs(tn.Left, target-tn.Val)
		dfs(tn.Right, target-tn.Val)
		path = path[:len(path)-1]
	}
	dfs(root, targetSum)
	return ans
}
