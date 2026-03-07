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

func inorderTraversal2(root *TreeNode) []int {
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

func inorderTraversal3(root *TreeNode) []int {
	ans := []int{}
	var dfs func(*TreeNode)
	dfs = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		dfs(tn.Left)
		ans = append(ans, tn.Val)
		dfs(tn.Right)
	}
	dfs(root)
	return ans
}

func inorderTraversal(root *TreeNode) []int {
	type colorNode struct {
		color int
		node  *TreeNode
	}
	stack := []colorNode{}
	GRAY, WITHE := 1, 0

	stack = append(stack, colorNode{WITHE, root})
	ans := []int{}
	for len(stack) > 0 {
		color, node := stack[len(stack)-1].color, stack[len(stack)-1].node
		stack = stack[:len(stack)-1]

		if node == nil {
			continue
		}

		if color == WITHE {
			stack = append(stack, colorNode{WITHE, node.Right})
			stack = append(stack, colorNode{GRAY, node})
			stack = append(stack, colorNode{WITHE, node.Left})
		} else {
			ans = append(ans, node.Val)
		}
	}
	return ans
}
