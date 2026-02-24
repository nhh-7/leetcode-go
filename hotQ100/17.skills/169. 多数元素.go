package skills

func majorityElement1(nums []int) int {
	votes, candidate := 0, 0
	for _, num := range nums {
		if votes == 0 {
			candidate = num
		}

		if candidate == num {
			votes++
		} else {
			votes--
		}
	}
	return candidate
}

func majorityElement(nums []int) int {
	vote := 0
	ans := nums[0]
	for _, num := range nums {
		if vote == 0 {
			ans = num
		}
		if num == ans {
			vote++
		} else {
			vote--
		}
	}
	return ans
}
