package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum1(root *TreeNode, targetSum int) int {
	cnt := make(map[int]int)
	cnt[0] = 1

	var traversal func(*TreeNode, int)
	var ans int
	traversal = func(tn *TreeNode, s int) {
		if tn == nil {
			return
		}
		s += tn.Val

		// 把 node 当作路径的终点，统计有多少个起点
		if c, ok := cnt[s-targetSum]; ok {
			ans += c
		}

		cnt[s]++

		traversal(tn.Left, s)
		traversal(tn.Right, s)

		cnt[s]--
	}
	traversal(root, 0)
	return ans
}

func pathSum(root *TreeNode, targetSum int) int {
	cnt := make(map[int]int)
	cnt[0] = 1
	ans := 0

	var traversal func(*TreeNode, int)
	traversal = func(tn *TreeNode, sum int) {
		if tn == nil {
			return
		}
		sum += tn.Val
		ans += cnt[sum-targetSum]
		cnt[sum]++

		traversal(tn.Left, sum)
		traversal(tn.Right, sum)

		cnt[sum]--
	}
	traversal(root, 0)
	return ans
}
