package times2

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func cutList(node *ListNode) *ListNode {
	slow, fast := node, node.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	p := slow.Next
	slow.Next = nil
	return p
}

func reorderList(head *ListNode) {
	l2 := cutList(head)
	var pre, cur *ListNode = nil, l2
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	l1, l2 := head, pre
	for l1 != nil && l2 != nil {
		l1Next := l1.Next
		l2Next := l2.Next

		l1.Next = l2
		l1 = l1Next

		l2.Next = l1Next
		l2 = l2Next
	}
}
