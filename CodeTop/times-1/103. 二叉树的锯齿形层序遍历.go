package times1

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	ans := [][]int{}
	isReversed := false
	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, size)
		for i := range size {
			top := queue[0]
			queue = queue[1:]
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}

			if !isReversed {
				level[i] = top.Val
			} else {
				level[size-i-1] = top.Val
			}
		}
		ans = append(ans, level)
		isReversed = !isReversed
	}
	return ans
}
