package hot

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	lenA, lenB := 0, 0
	curA, curB := headA, headB
	for curA != nil {
		lenA++
		curA = curA.Next
	}
	for curB != nil {
		lenB++
		curB = curB.Next
	}
	curA, curB = headA, headB
	gap := lenA - lenB
	if lenA < lenB {
		curA, curB = curB, curA
		gap = lenB - lenA
	}

	for gap > 0 {
		curA = curA.Next
		gap--
	}

	for curA != nil && curB != nil {
		if curA == curB {
			return curA
		}
		curA = curA.Next
		curB = curB.Next
	}
	return nil
}

/*
	思路：
	 - 设A链表长度为x，B链表长度为y，公共链表长度为z（z可以为0，表示不相交）
	 - 同时开始遍历两条链表， 知道节点相同
	 - 若不相同，且遍历到空节点那么跳转到另一条链表头部开始遍历
	 - 相同则退出遍历，最多遍历 x+y 次（也就是不相交）
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p, q := headA, headB

	for p != q {
		if p != nil {
			p = p.Next
		} else {
			p = headB
		}
		if q != nil {
			q = q.Next
		} else {
			q = headA
		}
	}
	return p
}
