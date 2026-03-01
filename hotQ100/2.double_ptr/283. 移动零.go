package hot

func moveZeroes1(nums []int) {
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

func moveZeroes(nums []int) {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
}
