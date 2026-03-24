package times2

import "container/heap"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type min_heap []*ListNode

func (m min_heap) Len() int {
	return len(m)
}

func (m min_heap) Less(i, j int) bool {
	return m[i].Val < m[j].Val
}

func (m min_heap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *min_heap) Push(v any) {
	*m = append(*m, v.(*ListNode))
}

func (m *min_heap) Pop() any {
	old := *m
	n := len(old)
	v := old[n-1]
	*m = (*m)[:n-1]
	return v
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := &min_heap{}
	for _, head := range lists {
		if head != nil {
			*h = append(*h, head)
		}
	}
	heap.Init(h)

	dummy := &ListNode{}
	cur := dummy

	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		cur.Next = node
		cur = cur.Next

		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}
	return dummy.Next
}
