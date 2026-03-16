package times1

import "math/rand"

func partition(nums []int, left, right int) int {
	idx := rand.Intn(right-left+1) + left
	nums[idx], nums[left] = nums[left], nums[idx]

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
	left, right := 0, len(nums)-1
	idx := len(nums) - k
	for left <= right {
		pivot := partition(nums, left, right)
		if pivot == idx {
			return nums[pivot]
		} else if pivot > idx {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}
	return -1
}
