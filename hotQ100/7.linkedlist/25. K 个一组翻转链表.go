package hot

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
*	- 参考92.反转链表2，这里固定反转长度为k
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	n := 0
	for cur := head; cur != nil; cur = cur.Next {
		n++
	}

	dummy := &ListNode{Next: head}

	p0 := dummy
	var pre, cur *ListNode = nil, head
	for i := k; i <= n; i += k {
		for j := 0; j < k; j++ {
			tmp := cur.Next
			cur.Next = pre
			pre = cur
			cur = tmp
		}
		p0.Next.Next = cur
		tmp := p0.Next
		p0.Next = pre
		p0 = tmp
	}
	return dummy.Next
}
