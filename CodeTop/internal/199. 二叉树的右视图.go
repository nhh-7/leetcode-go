package internal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		sz := len(queue)
		for i := range sz {
			tn := queue[0]
			queue = queue[1:]
			if tn.Left != nil {
				queue = append(queue, tn.Left)
			}
			if tn.Right != nil {
				queue = append(queue, tn.Right)
			}
			if i == sz-1 {
				ans = append(ans, tn.Val)
			}
		}
	}
	return ans
}
