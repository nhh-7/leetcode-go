package internal

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}

	p0 := dummy

	p1 := dummy.Next
	for p1 != nil && p1.Next != nil {
		p2 := p1.Next
		p3 := p1.Next.Next

		p0.Next = p2
		p2.Next = p1
		p1.Next = p3

		p0 = p1
		p1 = p3
	}
	return dummy.Next
}
