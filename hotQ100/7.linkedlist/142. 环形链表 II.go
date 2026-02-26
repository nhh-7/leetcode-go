package hot

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle1(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			cur := head
			for cur != slow {
				cur = cur.Next
				slow = slow.Next
			}
			return cur
		}
	}
	return nil
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if slow == fast {
			cur := head
			for cur != slow {
				cur = cur.Next
				slow = slow.Next
			}
			return cur
		}
	}
	return nil
}
