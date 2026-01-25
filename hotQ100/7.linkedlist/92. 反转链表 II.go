package hot

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
*
将要反转的独立出来，记录该段的前后节点，反转后在重新连接
*/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}

	p0 := dummy
	for i := 0; i < left-1; i++ {
		p0 = p0.Next
	}

	var pre, cur *ListNode = nil, p0.Next
	for k := 0; k < right-left+1; k++ {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	// 此时cur为该段后一节点，p0为该段前一节点，pre为该段最后一个节点

	p0.Next.Next = cur // 注意要先连接这里
	p0.Next = pre

	return dummy.Next
}
