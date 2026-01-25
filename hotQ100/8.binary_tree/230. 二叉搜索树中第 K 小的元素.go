package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest1(root *TreeNode, k int) int {
	var ans int
	var traversal func(*TreeNode) bool
	traversal = func(tn *TreeNode) bool {
		if tn == nil {
			return false
		}
		if traversal(tn.Left) {
			return true
		}
		k--
		if k == 0 {
			ans = tn.Val
			return true
		}
		if traversal(tn.Right) {
			return true
		}
		return false
	}
	traversal(root)
	return ans
}

type colorNode struct {
	color int // 0-白色，1-灰色
	node  *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	stk := []colorNode{}
	stk = append(stk, colorNode{color: 0, node: root})

	for len(stk) > 0 {
		cur := stk[len(stk)-1]
		stk = stk[:len(stk)-1]

		if cur.node == nil {
			continue
		}
		if cur.color == 0 {
			stk = append(stk, colorNode{color: 0, node: cur.node.Right})
			stk = append(stk, colorNode{color: 1, node: cur.node})
			stk = append(stk, colorNode{color: 0, node: cur.node.Left})
		} else {
			k--
			if k == 0 {
				return cur.node.Val
			}
		}
	}
	return 0
}
