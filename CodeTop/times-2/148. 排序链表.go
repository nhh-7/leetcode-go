package times2

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func cut(node *ListNode) *ListNode {
	slow, fast := node, node.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	p := slow.Next
	slow.Next = nil
	return p
}

func mergeList(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return dummy.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := cut(head)
	head = sortList(head)
	mid = sortList(mid)
	return mergeList(head, mid)
}
