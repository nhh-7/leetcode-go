package times2

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	p0 := dummy
	for i := 0; i < left-1; i++ {
		p0 = p0.Next
	}
	var pre, cur *ListNode = nil, p0.Next
	for i := 0; i < right-left+1; i++ {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	p0.Next.Next = cur
	p0.Next = pre
	return dummy.Next
}
