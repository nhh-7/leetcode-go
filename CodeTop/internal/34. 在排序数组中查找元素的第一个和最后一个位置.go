package internal

func getFirsh(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (right-left)/2 + left

		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func getLast(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (right-left)/2 + left

		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}

func searchRange(nums []int, target int) []int {
	left, right := getFirsh(nums, target), getLast(nums, target)

	if left >= 0 && left < len(nums) && nums[left] == target {
		return []int{left, right}
	}
	return []int{-1, -1}
}
