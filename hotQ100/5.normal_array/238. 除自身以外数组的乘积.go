package hot

func productExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)

	prefix := make([]int, n)
	prefix[0] = 1
	for i := 1; i < n; i++ {
		prefix[i] = nums[i-1] * prefix[i-1]
	}

	sufix := make([]int, n)
	sufix[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		sufix[i] = nums[i+1] * sufix[i+1]
	}

	for idx, num := range prefix {
		answer[idx] = num * sufix[idx]
	}
	return answer
}
