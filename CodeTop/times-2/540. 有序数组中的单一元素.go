package times2

func singleNonDuplicate(nums []int) int {
	n := len(nums)
	left, right := 0, n-1
	for left < right {
		mid := left + (right-left)/2
		// 要保证mid一定是偶数，这样每次都比较 （偶，奇）这个序列

		// 如果 mid 是偶数，则 mid^1 是 mid+1
		// 如果 mid 是奇数，则 mid^1 是 mid-1
		if nums[mid] == nums[mid^1] { // 说明下标顺序没有被打乱，唯一出现的数在右边
			left = mid + 1
		} else { // 说明下标顺序没有被打乱，唯一出现的数在左边或当前这个数
			right = mid
		}
	}
	return left
}
