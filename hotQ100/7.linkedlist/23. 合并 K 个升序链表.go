package hot

import "container/heap"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// type minHeap []*ListNode

// func (h *minHeap) Push(v any) {
// 	*h = append(*h, v.(*ListNode))
// }
// func (h *minHeap) Pop() any {
// 	old := *h
// 	n := len(old)
// 	v := old[n-1]
// 	*h = old[:n-1]
// 	return v
// }
// func (h minHeap) Len() int {
// 	return len(h)
// }
// func (h minHeap) Less(i, j int) bool {
// 	return h[i].Val < h[j].Val
// }
// func (h minHeap) Swap(i, j int) {
// 	h[i], h[j] = h[j], h[i]
// }

// func mergeKLists(lists []*ListNode) *ListNode {
// 	mh := minHeap{}
// 	for _, list := range lists {
// 		if list != nil {
// 			mh = append(mh, list)
// 		}
// 	}
// 	heap.Init(&mh)

// 	dummy := &ListNode{}
// 	cur := dummy

// 	for mh.Len() > 0 {
// 		node := heap.Pop(&mh).(*ListNode)
// 		if node.Next != nil {
// 			heap.Push(&mh, node.Next)
// 		}
// 		cur.Next = node
// 		cur = cur.Next
// 	}
// 	return dummy.Next
// }

type min_heap []*ListNode

func (m min_heap) Less(i, j int) bool {
	return m[i].Val < m[j].Val
}
func (m min_heap) Len() int {
	return len(m)
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
	old = old[:n-1]
	*m = old
	return v
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := &min_heap{}
	for _, node := range lists {
		if node != nil {
			h.Push(node)
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
	cur.Next = nil
	return dummy.Next
}
