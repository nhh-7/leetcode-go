package hot

func findDisappearedNumbers1(nums []int) []int {
	for _, num := range nums {
		val := abs(num)
		idx := val - 1
		if nums[idx] > 0 {
			nums[idx] *= -1
		}
	}

	ans := []int{}
	for i, num := range nums {
		if num > 0 {
			ans = append(ans, i+1)
		}
	}
	return ans
}

func findDisappearedNumbers(nums []int) []int {
	for _, v := range nums {
		if nums[abs(v)-1] > 0 {
			nums[abs(v)-1] *= -1
		}
	}
	ans := []int{}
	for i, v := range nums {
		if v > 0 {
			ans = append(ans, i+1)
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
