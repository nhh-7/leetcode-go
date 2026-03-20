package times2

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	n := 0
	for cur := head; cur != nil; cur = cur.Next {
		n++
	}

	dummy := &ListNode{Next: head}
	p0 := dummy

	for i := k; i <= n; i += k {
		var pre, cur *ListNode = nil, p0.Next
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
