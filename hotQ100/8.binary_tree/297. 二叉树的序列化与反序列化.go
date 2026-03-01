package binarytree

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	res := []string{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node != nil {
			res = append(res, strconv.Itoa(node.Val))
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		} else {
			res = append(res, "null")
		}
	}
	return strings.Join(res, ",")
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	vals := strings.Split(data, ",")
	val, _ := strconv.Atoi(vals[0])
	root := &TreeNode{Val: val}
	queue := []*TreeNode{root}
	cursor := 1

	for len(queue) > 0 && cursor < len(vals) {
		parent := queue[0]
		queue = queue[1:]

		if vals[cursor] != "null" {
			lv, _ := strconv.Atoi(vals[cursor])
			parent.Left = &TreeNode{Val: lv}
			queue = append(queue, parent.Left)
		}
		cursor++

		if cursor < len(vals) && vals[cursor] != "null" {
			rv, _ := strconv.Atoi(vals[cursor])
			parent.Right = &TreeNode{Val: rv}
			queue = append(queue, parent.Right)
		}
		cursor++
	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
