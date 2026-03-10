package internal

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	queue := []*TreeNode{}
	ans := [][]int{}
	if root == nil {
		return ans
	}

	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, size)
		for i := range size {
			node := queue[0]
			queue = queue[1:]
			level[i] = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, level)
	}
	return ans
}
