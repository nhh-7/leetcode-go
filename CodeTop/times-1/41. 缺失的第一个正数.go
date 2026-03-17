package times1

func firstMissingPositive(nums []int) int {
	n := len(nums)

	for i := range nums {
		for nums[i] != i+1 {
			if nums[i] <= 0 || nums[i] > n || nums[i] == nums[nums[i]-1] {
				break
			}
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	for i, v := range nums {
		if v != i+1 {
			return i + 1
		}
	}
	return n + 1
}
