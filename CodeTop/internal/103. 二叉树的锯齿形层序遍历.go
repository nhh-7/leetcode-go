package internal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	queue := []*TreeNode{}
	if root == nil {
		return ans
	}
	queue = append(queue, root)
	isReverse := false
	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, size)
		for i := range size {
			node := queue[0]
			queue = queue[1:]
			if !isReverse {
				level[i] = node.Val
			} else {
				level[size-i-1] = node.Val // 逆序
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, level)
		isReverse = !isReverse
	}
	return ans
}
