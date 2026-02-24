package heap

import (
	"container/heap"
	"math/rand"
)

type minHeap []int

func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h minHeap) Len() int {
	return len(h)
}

func (h *minHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func findKthLargest1(nums []int, k int) int {
	h := &minHeap{}
	heap.Init(h)
	for _, num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return (*h)[0]
}

// 不需要快排完全排序，只需要第k大的元素到达了位置上，那么其右边的元素肯定是比他大的，左边的肯定是比他小的
func partition1(nums []int, left, right int) int {
	random_idx := left + rand.Intn(right-left+1)
	nums[left], nums[random_idx] = nums[random_idx], nums[left]
	pivot := nums[left]

	i, j := left+1, right
	for {
		for i <= j && nums[i] < pivot {
			i++
		}
		for i <= j && nums[j] > pivot {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	nums[left], nums[j] = nums[j], nums[left]
	return j
}

func findKthLargest2(nums []int, k int) int {
	// 寻找在最终有序序列中 n-k 下标的元素
	n := len(nums)
	left, right := 0, n-1
	for {
		pivot := partition1(nums, left, right)
		if pivot == n-k {
			return nums[pivot]
		} else if pivot > n-k {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}
}

func partition(nums []int, left, right int) int {
	pivot := nums[left]
	i, j := left+1, right

	for {
		for i <= j && nums[i] < pivot {
			i++
		}
		for i <= j && nums[j] > pivot {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	nums[left], nums[j] = nums[j], nums[left]
	return j
}

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	left, right := 0, n-1
	for {
		pivot := partition(nums, left, right)
		if pivot == n-k {
			return nums[pivot]
		} else if pivot > n-k {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}
}
