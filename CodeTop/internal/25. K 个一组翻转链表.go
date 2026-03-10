package internal

func reverseKGroup(head *ListNode, k int) *ListNode {
	n := 0
	cnt := head
	for cnt != nil {
		n++
		cnt = cnt.Next
	}

	dummy := &ListNode{Next: head}

	p0 := dummy // p0是 翻转每段的前一节点
	var pre, cur *ListNode = nil, p0.Next
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
