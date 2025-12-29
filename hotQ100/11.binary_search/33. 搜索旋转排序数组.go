package binarysearch

//
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		// mid 只会处于一段有序区间中，也就是说，要么[left, mid]有序，要么[mid, right]有序
		if nums[mid] >= nums[left] {
			if nums[mid] > target && nums[left] <= target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && nums[right] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
