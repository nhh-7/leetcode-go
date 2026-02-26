package hot

func findDisappearedNumbers(nums []int) []int {
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
