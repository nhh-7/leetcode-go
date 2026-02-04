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
			// 说明 mid 到 right 这一段是升序的，最小值一定不在 mid 的右边（但 mid 本身可能是最小值）。因此，收缩右边界：right = mid。
			if right == mid {
				return nums[right]
			}
			right = mid
		}
	}
	return nums[right]
}
