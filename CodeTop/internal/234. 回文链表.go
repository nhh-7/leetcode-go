package internal

func searchMid(node *ListNode) *ListNode { // 必须要node不为nil
	if node == nil {
		return nil
	}
	slow, fast := node, node.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	p := slow.Next
	slow.Next = nil
	return p
}

func isPalindrome(head *ListNode) bool {
	mid := searchMid(head)

	var pre, cur *ListNode = nil, mid
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	for head != nil && pre != nil {
		if head.Val != pre.Val {
			return false
		}
		head = head.Next
		pre = pre.Next
	}
	return true
}
