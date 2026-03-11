package internal

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func findMid(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reorderList(head *ListNode) {
	mid := findMid(head)
	l1, l2 := head, mid.Next
	mid.Next = nil

	var pre, cur *ListNode = nil, l2
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	l2 = pre

	for l1 != nil && l2 != nil {
		l1Next, l2Next := l1.Next, l2.Next

		l1.Next = l2
		l1 = l1Next

		l2.Next = l1
		l2 = l2Next
	}
}
