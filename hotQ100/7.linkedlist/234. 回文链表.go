package hot

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func isPalindrome(head *ListNode) bool {
	midNode := middleNode(head)
	var pre, cur *ListNode = nil, midNode
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	cur1, cur2 := head, pre
	for cur1 != nil && cur2 != nil {
		if cur1.Val != cur2.Val {
			return false
		}
		cur1, cur2 = cur1.Next, cur2.Next
	}
	return true
}
