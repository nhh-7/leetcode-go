package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	var traversal func(*TreeNode) int

	/*
		traversal的作用是返回node到其叶子节点最长长度
		traversal(root)时便将所有节点访问一遍，则所有路径都尝试了一遍，用ans记录最大值
	*/
	traversal = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := traversal(node.Left)
		right := traversal(node.Right)
		ans = max(ans, left+right+1)
		return max(left, right) + 1
	}
	traversal(root)
	return ans - 1
}
