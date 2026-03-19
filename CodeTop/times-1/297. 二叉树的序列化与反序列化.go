package times1

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
	res := []string{}
	if root == nil {
		return ""
	}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		if top == nil {
			res = append(res, "null")
		} else {
			res = append(res, strconv.Itoa(top.Val))
			queue = append(queue, top.Left)
			queue = append(queue, top.Right)
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

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
