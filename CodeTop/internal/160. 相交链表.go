package internal

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	cur_a, cur_b := headA, headB
	for cur_a != cur_b {
		if cur_a == nil {
			cur_a = headB
		} else {
			cur_a = cur_a.Next
		}
		if cur_b == nil {
			cur_b = headA
		} else {
			cur_b = cur_b.Next
		}
	}
	return cur_a
}
