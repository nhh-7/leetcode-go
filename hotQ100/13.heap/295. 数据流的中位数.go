package heap

import "container/heap"

// left为最大堆，right为最小堆，分别存储数据流的有序状态下的左右两部分
type MedianFinder struct {
	left, right hp
}

// 使用同一个最小堆实现最大最小堆，将数据放入left最大堆时，对数据取反，取出要使用再把数据变换回来，这样则实现了最大堆

func Constructor() MedianFinder {
	return MedianFinder{}
}

// 规定left长度大于或等于right的长度，大于也只可能大一个
func (mf *MedianFinder) AddNum(num int) {
	if mf.left.Len() == mf.right.Len() {
		heap.Push(&mf.right, num)
		heap.Push(&mf.left, -heap.Pop(&mf.right).(int))
	} else if mf.left.Len() > mf.right.Len() {
		heap.Push(&mf.left, -num)
		heap.Push(&mf.right, -heap.Pop(&mf.left).(int))
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.left.Len() > mf.right.Len() {
		return float64(-mf.left[0])
	} else {
		return float64((mf.right[0] - mf.left[0]) / 2.0)
	}
}

type hp []int

func (h hp) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h hp) Len() int {
	return len(h)
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(v any) {
	*h = append(*h, v.(int))
}

func (h *hp) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	old = old[:n-1]
	*h = old
	return x
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
