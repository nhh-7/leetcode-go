package binarysearch

func getLast1(nums []int, target int) int {
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

func getFirst1(nums []int, target int) int {
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
func searchRange1(nums []int, target int) []int {
	fst, lst := getFirst1(nums, target), getLast1(nums, target)

	if fst >= 0 && fst < len(nums) && nums[fst] == target {
		return []int{fst, lst}
	}
	return []int{-1, -1}
}

func getFirst(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] == target { // 继续向左找
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	// 最后时， right落在结束位置的下一个位置
	return left
}

func getLast(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] == target {
			// 继续向右找
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	// 最后时， left落在开始位置的上一个位置
	return right
}

func searchRange(nums []int, target int) []int {
	fst, lst := getFirst(nums, target), getLast(nums, target)
	if fst >= 0 && fst < len(nums) && nums[fst] == target {
		return []int{fst, lst}
	}
	return []int{-1, -1}
}
