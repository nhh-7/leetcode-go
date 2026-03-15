package internal

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func getMid(node *ListNode) *ListNode {
	slow, fast := node, node
	pre := slow
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		pre = slow
		slow = slow.Next
	}
	pre.Next = nil
	return slow
}

func mergeList(node1, node2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	for node1 != nil && node2 != nil {
		if node1.Val < node2.Val {
			cur.Next = node1
			node1 = node1.Next
		} else {
			cur.Next = node2
			node2 = node2.Next
		}
		cur = cur.Next
	}
	if node1 != nil {
		cur.Next = node1
	}
	if node2 != nil {
		cur.Next = node2
	}
	return dummy.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := getMid(head)

	head = sortList(head)
	mid = sortList(mid)

	return mergeList(head, mid)
}
