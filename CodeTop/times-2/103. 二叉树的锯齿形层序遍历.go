package times2

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := [][]int{}

	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}
	isReversed := false

	for len(queue) > 0 {
		sz := len(queue)
		level := make([]int, sz)

		for i := 0; i < sz; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			if !isReversed {
				level[i] = node.Val
			} else {
				level[sz-i-1] = node.Val
			}
		}
		isReversed = !isReversed
		ans = append(ans, level)
	}
	return ans
}
