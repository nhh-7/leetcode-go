package hot

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// func cutList(head *ListNode) *ListNode {
// 	fast, slow := head, head
// 	pre := head
// 	for fast != nil && fast.Next != nil {
// 		pre = slow
// 		slow = slow.Next
// 		fast = fast.Next.Next
// 	}
// 	pre.Next = nil
// 	return slow
// }

// func merge(h1, h2 *ListNode) *ListNode {
// 	dummy := &ListNode{}
// 	cur := dummy
// 	for h1 != nil && h2 != nil {
// 		if h1.Val < h2.Val {
// 			cur.Next = h1
// 			h1 = h1.Next
// 		} else {
// 			cur.Next = h2
// 			h2 = h2.Next
// 		}
// 		cur = cur.Next
// 	}
// 	if h1 != nil {
// 		cur.Next = h1
// 	} else {
// 		cur.Next = h2
// 	}
// 	return dummy.Next
// }

// func sortList(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	mid := cutList(head)

// 	head = sortList(head)
// 	mid = sortList(mid)
// 	return merge(head, mid)
// }

func cutList(head *ListNode) *ListNode {
	slow, fast := head, head
	pre := slow
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	pre.Next = nil
	return slow
}

func merge(h1, h2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for h1 != nil && h2 != nil {
		if h1.Val < h2.Val {
			cur.Next = h1
			h1 = h1.Next
		} else {
			cur.Next = h2
			h2 = h2.Next
		}
		cur = cur.Next
	}
	if h1 != nil {
		cur.Next = h1
	}
	if h2 != nil {
		cur.Next = h2
	}
	return dummy.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := cutList(head)

	head = sortList(head)
	mid = sortList(mid)
	return merge(mid, head)
}
