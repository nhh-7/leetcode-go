package times1

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func cutList(ln *ListNode) *ListNode {
	slow, fast := ln, ln.Next // 偶数节点数时， 分为slow指向偏右或偏左两种写法
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	l := slow.Next
	slow.Next = nil
	return l
}

func mergeList(ln1, ln2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	for ln1 != nil && ln2 != nil {
		if ln1.Val < ln2.Val {
			cur.Next = ln1
			ln1 = ln1.Next
		} else {
			cur.Next = ln2
			ln2 = ln2.Next
		}
		cur = cur.Next
	}
	if ln1 != nil {
		cur.Next = ln1
	}
	if ln2 != nil {
		cur.Next = ln2
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
	return mergeList(head, mid)
}
