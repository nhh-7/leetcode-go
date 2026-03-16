package times1

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	cnt := 0
	for cur := head; cur != nil; cur = cur.Next {
		cnt++
	}

	Ppre := dummy
	var pre, cur *ListNode = nil, head
	for i := k; i <= cnt; i += k {
		for j := 0; j < k; j++ {
			tmp := cur.Next
			cur.Next = pre
			pre = cur
			cur = tmp
		}
		Ppre.Next.Next = cur
		tmp := Ppre.Next
		Ppre.Next = pre
		pre = tmp
		Ppre = pre
	}
	return dummy.Next
}
