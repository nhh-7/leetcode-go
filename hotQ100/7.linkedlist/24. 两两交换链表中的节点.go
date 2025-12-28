package hot

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**
 0 -> 1 -> 2 -> 3 (0为虚拟节点)
三步走：
 - 0 -> 2
 - 2 -> 1
 - 1 -> 3
迭代更新node0和node1
*/
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	node0 := dummy // 被交换两节点的前一节点
	node1 := head
	for node1 != nil && node1.Next != nil {
		node2 := node1.Next
		node3 := node1.Next.Next

		node0.Next = node2
		node2.Next = node1
		node1.Next = node3

		node0 = node1
		node1 = node3
	}
	return dummy.Next
}

/**
 0 -> 1 -> 2 -> 3 (0为虚拟节点)
三步走：
 - 0 -> 2
 - 2 -> 1
 - 1 -> 3
*/
func swapPairs1(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	cur := dummy // 被交换两节点的前一节点
	for cur.Next != nil && cur.Next.Next != nil {
		tmp1 := cur.Next         // 被交换的第一节点
		cur.Next = cur.Next.Next // 前一节点 -> 被交换的第二节点

		tmp2 := cur.Next.Next // 被交换两节点的后一节点
		cur.Next.Next = tmp1  // 第二节点 -> 第一节点

		cur.Next.Next.Next = tmp2 // 连接后一节点
		cur = cur.Next.Next       // 更新前一节点
	}
	return dummy.Next
}
