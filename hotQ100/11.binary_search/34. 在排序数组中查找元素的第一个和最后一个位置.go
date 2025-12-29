package binarysearch

func getLast(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}

func getFirst(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

// 使用两次二分查找，分别查找目标值的上界和下界
func searchRange(nums []int, target int) []int {
	fst, lst := getFirst(nums, target), getLast(nums, target)

	if fst >= 0 && fst < len(nums) && nums[fst] == target {
		return []int{fst, lst}
	}
	return []int{-1, -1}
}
