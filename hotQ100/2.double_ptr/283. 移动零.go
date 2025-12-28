package hot

func moveZeroes(nums []int) {
	left, right := 0, 0
	len := len(nums)

	for ; right < len; right++ {
		if nums[right] == 0 {
			continue
		}
		nums[left] = nums[right]
		left++
	}

	for left < len {
		nums[left] = 0
		left++
	}
}
