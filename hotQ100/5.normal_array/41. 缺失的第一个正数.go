package hot

func firstMissingPositive1(nums []int) int {
	n := len(nums)

	for i := 0; i < n; i++ {
		for nums[i] != i+1 {
			if nums[i] <= 0 || nums[i] > n || nums[i] == nums[nums[i]-1] {
				break
			}
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i, num := range nums {
		if num != i+1 {
			return i + 1
		}
	}
	return n + 1
}

func firstMissingPositive(nums []int) int {
	n := len(nums)

	for i := range nums {
		for nums[i] != i+1 {
			if nums[i] <= 0 || nums[i] > n || nums[i] == nums[nums[i]-1] {
				break
			}
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i, x := range nums {
		if x != i+1 {
			return i + 1
		}
	}
	return n + 1
}
