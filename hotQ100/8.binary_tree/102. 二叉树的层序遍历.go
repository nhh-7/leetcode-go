package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	q := []*TreeNode{}
	ans := [][]int{}

	q = append(q, root)

	for len(q) > 0 {
		n := len(q)
		layer := make([]int, n)
		for i := range n {
			node := q[0]
			q = q[1:]
			layer[i] = node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		ans = append(ans, layer)
	}
	return ans
}
