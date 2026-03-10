package internal

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

// 找最终下标在 n-k 位置的元素
func findKthLargest(nums []int, k int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		pivot := partition(nums, left, right)
		if pivot == n-k {
			return nums[pivot]
		} else if pivot > n-k {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}
	return -1
}
