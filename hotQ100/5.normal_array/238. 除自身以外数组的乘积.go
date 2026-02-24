package hot

func productExceptSelf1(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)

	prefix := make([]int, n)
	prefix[0] = 1
	for i := 1; i < n; i++ {
		prefix[i] = nums[i-1] * prefix[i-1]
	}
	// prefix[i]存储 从nums[0] 到 nums[i-1]的乘积
	// sufix[i]存储 从nums[i+1] 到 nums[n-1]的乘积

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

func productExceptSelf(nums []int) []int {
	// prefix[i]为nums[i]之前的所有元素的乘积
	// sufix[i]为nums[i]之后所有元素的乘积
	n := len(nums)
	prefix, sufix := make([]int, n), make([]int, n)
	prefix[0] = 1
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}
	sufix[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		sufix[i] = sufix[i+1] * nums[i+1]
	}
	ans := make([]int, n)
	for i := range ans {
		ans[i] = prefix[i] * sufix[i]
	}
	return ans
}
