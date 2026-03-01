package heap

import "container/heap"

// type pair struct {
// 	v, q int
// }

// type min_heap []pair

// func (mh min_heap) Less(i, j int) bool {
// 	return mh[i].q < mh[j].q
// }

// func (mh min_heap) Len() int {
// 	return len(mh)
// }

// func (mh min_heap) Swap(i, j int) {
// 	mh[i], mh[j] = mh[j], mh[i]
// }

// func (mh *min_heap) Push(v any) {
// 	*mh = append(*mh, v.(pair))
// }

// func (mh *min_heap) Pop() any {
// 	old := *mh
// 	n := len(old)
// 	v := old[n-1]
// 	*mh = old[:n-1] // 这里删除的最后一位是已经被交换到数组末尾的堆顶元素
// 	return v
// }

// func topKFrequent1(nums []int, k int) []int {
// 	h := &min_heap{}
// 	heap.Init(h)
// 	m := map[int]int{}
// 	for _, nums := range nums {
// 		m[nums]++
// 	}

// 	for v, q := range m {
// 		heap.Push(h, pair{v, q})
// 		if h.Len() > k {
// 			heap.Pop(h)
// 		}
// 	}
// 	ans := []int{}
// 	for h.Len() > 0 {
// 		ans = append(ans, heap.Pop(h).(pair).v)
// 	}
// 	return ans
// }

// 不使用堆
func topKFrequent2(nums []int, k int) []int {
	// 第一步：统计每个元素的出现次数
	cnt := make(map[int]int)
	maxCnt := 0
	for _, num := range nums {
		cnt[num]++
		maxCnt = max(maxCnt, cnt[num])
	}

	// 第二步：把出现次数相同的元素，放到同一个桶中
	buckets := make([][]int, maxCnt+1)
	for v, q := range cnt {
		buckets[q] = append(buckets[q], v)
	}

	// 第三步：倒序遍历 buckets，把出现次数前 k 大的元素加入答案
	ans := make([]int, 0, k)
	for i := maxCnt; i >= 0 && len(ans) < k; i-- {
		ans = append(ans, buckets[i]...)
	}
	return ans
}

func topKFrequent3(nums []int, k int) []int {
	cnt := make(map[int]int)
	maxCnt := 0
	for _, n := range nums {
		cnt[n]++
		maxCnt = max(maxCnt, cnt[n])
	}

	buckets := make([][]int, maxCnt+1)

	for k, v := range cnt {
		buckets[v] = append(buckets[v], k)
	}

	ans := make([]int, 0, k)
	for i := maxCnt; i >= 0 && len(ans) < k; i-- {
		ans = append(ans, buckets[i]...)
	}
	return ans
}

type pair struct {
	num, q int
}

type min_heap []pair

func (mh min_heap) Less(i, j int) bool {
	return mh[i].q < mh[j].q
}

func (mh min_heap) Len() int {
	return len(mh)
}

func (mh min_heap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *min_heap) Push(v any) {
	*mh = append(*mh, v.(pair))
}

func (mh *min_heap) Pop() any {
	n := len(*mh)
	v := (*mh)[n-1]
	*mh = (*mh)[:n-1]
	return v
}

func topKFrequent(nums []int, k int) []int {
	h := &min_heap{}
	heap.Init(h)
	cnt := make(map[int]int)
	for _, num := range nums {
		cnt[num]++
	}
	for num, q := range cnt {
		heap.Push(h, pair{num, q})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ans := make([]int, 0, k)
	for h.Len() > 0 {
		ans = append(ans, heap.Pop(h).(pair).num)
	}
	return ans
}
