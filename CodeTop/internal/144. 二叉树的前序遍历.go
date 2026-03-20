package internal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal1(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	var dfs func(tn *TreeNode)
	dfs = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		ans = append(ans, tn.Val)
		dfs(tn.Left)
		dfs(tn.Right)
	}
	dfs(root)
	return ans
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	type pair struct {
		color int
		tn    *TreeNode
	}

	ans := []int{}
	if root == nil {
		return ans
	}
	WHITE, GRAY := 0, 1
	stk := []pair{{0, root}}
	for len(stk) > 0 {
		p := stk[len(stk)-1]
		stk = stk[:len(stk)-1]

		if p.tn == nil {
			continue
		}

		if p.color == WHITE {
			stk = append(stk, pair{WHITE, p.tn.Right})
			stk = append(stk, pair{WHITE, p.tn.Left})
			stk = append(stk, pair{GRAY, p.tn})
		} else {
			ans = append(ans, p.tn.Val)
		}
	}
	return ans
}
