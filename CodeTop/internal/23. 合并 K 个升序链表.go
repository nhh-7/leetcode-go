package internal

import (
	"container/heap"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type min_heap []*ListNode

func (mh min_heap) Len() int {
	return len(mh)
}

func (mh min_heap) Less(i, j int) bool {
	return mh[i].Val < mh[j].Val
}

func (mh min_heap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *min_heap) Push(v any) {
	*mh = append(*mh, v.(*ListNode))
}

func (mh *min_heap) Pop() any {
	old := *mh
	n := len(old)
	v := old[n-1]
	*mh = (*mh)[:n-1]
	return v
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := &min_heap{}
	heap.Init(h)

	for _, head := range lists {
		if head != nil {
			heap.Push(h, head)
		}
	}

	dummy := &ListNode{}
	cur := dummy
	for h.Len() > 0 {
		cur.Next = heap.Pop(h).(*ListNode)
		cur = cur.Next
		if cur.Next != nil {
			heap.Push(h, cur.Next)
		}
	}
	return dummy.Next
}
