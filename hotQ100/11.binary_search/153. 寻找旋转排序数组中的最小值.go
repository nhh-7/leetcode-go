package binarysearch

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		// 判断数组是否被分为了两段，然后判断mid处于哪一段
		if nums[mid] > nums[len(nums)-1] {
			// 表示分为了两段有序，前一段大于后一段，且mid处于前一段, 最小值肯定在后一段
			left = mid + 1
		} else {
			// 表示从mid到最后是有序的，最小值在 mid 或 mid之前
			if right == mid {
				return nums[right]
			}
			right = mid
		}
	}
	return nums[right]
}
