package internal

import "math/rand"

func partition1(nums []int, left, right int) int {
	r := rand.Intn(right-left+1) + left
	nums[left], nums[r] = nums[r], nums[left]

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
	nums[j], nums[left] = nums[left], nums[j]
	return j
}

func QuickSort(nums []int, left, right int) {
	if left < right {
		pivot := partition1(nums, left, right)

		QuickSort(nums, left, pivot-1)
		QuickSort(nums, pivot+1, right)
	}
}

func sortArray(nums []int) []int {
	QuickSort(nums, 0, len(nums)-1)
	return nums
}
