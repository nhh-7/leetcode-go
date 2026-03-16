package internal

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func trainingPlan(head *ListNode, cnt int) *ListNode {
	slow, fast := head, head

	for cnt > 0 {
		fast = fast.Next
		cnt--
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
