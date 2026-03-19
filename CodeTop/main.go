package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 输入有序序列，直接构建二叉搜索树
func buildTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = buildTree(nums[:mid])
	root.Right = buildTree(nums[mid+1:])
	return root
}

// 给定一个层序遍历的字符串数组，空节点对应null
func buildTree1(vals []string) *TreeNode {
	if len(vals) == 0 {
		return nil
	}
	v, _ := strconv.Atoi(vals[0])
	root := &TreeNode{Val: v}
	queue := []*TreeNode{root}

	idx := 1
	for len(queue) > 0 && idx < len(vals) {
		parent := queue[0]
		queue = queue[1:]

		if idx < len(vals) && vals[idx] != "null" {
			v, _ := strconv.Atoi(vals[idx])
			parent.Left = &TreeNode{Val: v}
			queue = append(queue, parent.Left)
		}
		idx++

		if idx < len(vals) && vals[idx] != "null" {
			v, _ := strconv.Atoi(vals[idx])
			parent.Right = &TreeNode{Val: v}
			queue = append(queue, parent.Right)
		}
		idx++
	}
	return root
}

func preorder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	preorder(root.Left)
	preorder(root.Right)
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	fmt.Printf("%d ", root.Val)
	inorder(root.Right)
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	vals := []string{"1", "2", "3", "null", "null", "4", "5"}

	root1 := buildTree(nums)
	root2 := buildTree1(vals)

	preorder(root1)
	fmt.Println()
	inorder(root1)
	fmt.Println()

	preorder(root2)
	fmt.Println()
	inorder(root2)
	fmt.Println()
}
