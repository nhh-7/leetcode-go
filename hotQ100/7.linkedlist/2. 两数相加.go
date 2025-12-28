package hot

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	j := 0
	for l1 != nil && l2 != nil {
		val := (l1.Val + l2.Val + j) % 10
		j = (l1.Val + l2.Val + j) / 10
		l1.Val = val
		cur.Next = l1
		l1 = l1.Next
		l2 = l2.Next
		cur = cur.Next
	}
	fmt.Print(dummy)
	for l1 != nil {
		val := (l1.Val + j) % 10
		j = (l1.Val + j) / 10
		l1.Val = val
		cur.Next = l1
		l1 = l1.Next
		cur = cur.Next
	}
	for l2 != nil {
		val := (l2.Val + j) % 10
		j = (l2.Val + j) / 10
		l2.Val = val
		cur.Next = l2
		l2 = l2.Next
		cur = cur.Next
	}
	if j != 0 {
		cur.Next = &ListNode{
			Val:  j,
			Next: nil,
		}
	}
	return dummy.Next
}
