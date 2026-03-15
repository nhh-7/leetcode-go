package internal

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast := head, head
	for n > 0 {
		fast = fast.Next
		n--
	}
	dummy := &ListNode{Next: head}
	pre := dummy
	for fast != nil {
		fast = fast.Next
		pre = slow
		slow = slow.Next
	}
	pre.Next = pre.Next.Next
	return dummy.Next
}
