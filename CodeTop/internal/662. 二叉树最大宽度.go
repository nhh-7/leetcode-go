package internal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	type pair struct {
		node *TreeNode
		idx  int
	}
	queue := []pair{{root, 1}}

	maxWidth := 0
	for len(queue) > 0 {
		size := len(queue)
		startIdx := queue[0].idx // 每一层的起点编号

		for i := 0; i < size; i++ {
			p := queue[0]
			queue = queue[1:]

			node, currIdx := p.node, p.idx
			// 更新当前层宽度（每次循环都会更新，最后一次即为 end - start + 1）
			if currIdx-startIdx+1 > maxWidth {
				maxWidth = currIdx - startIdx + 1
			}

			// 关键：为了防止编号过大溢出，我们可以将编号减去当前层起点偏移量
			// 这样每一层的编号都会重新从相对较小的值开始计算
			relativeIdx := currIdx - startIdx

			if node.Left != nil {
				queue = append(queue, pair{node.Left, relativeIdx * 2})
			}
			if node.Right != nil {
				queue = append(queue, pair{node.Right, relativeIdx*2 + 1})
			}
		}
	}
	return maxWidth
}
