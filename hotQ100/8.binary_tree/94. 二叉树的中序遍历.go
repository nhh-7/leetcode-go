package binarytree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal1(root *TreeNode) (ans []int) {
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		ans = append(ans, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return ans
}

type ColorAndNode struct {
	color int
	node  *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	stack := []ColorAndNode{}
	WHITE, GRAY := 0, 1
	stack = append(stack, ColorAndNode{color: WHITE, node: root})

	ans := []int{}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if top.node == nil {
			continue
		}

		if top.color == WHITE {
			stack = append(stack, ColorAndNode{
				color: WHITE,
				node:  top.node.Right,
			})
			stack = append(stack, ColorAndNode{
				color: GRAY,
				node:  top.node,
			})
			stack = append(stack, ColorAndNode{
				color: WHITE,
				node:  top.node.Left,
			})
		} else {
			ans = append(ans, top.node.Val)
		}
	}
	return ans
}
