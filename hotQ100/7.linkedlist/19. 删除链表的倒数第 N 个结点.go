package hot

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	slow, fast := dummy, dummy
	for n > 0 {
		fast = fast.Next
		n--
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	step := n + 1
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy

	for step > 0 && fast != nil {
		fast = fast.Next
		step--
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
