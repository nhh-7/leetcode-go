package hot

func maxSlidingWindow(nums []int, k int) []int {
	ans := make([]int, len(nums)-k+1)
	q := []int{} // 模拟单调队列

	for i, x := range nums {
		for len(q) > 0 && nums[q[len(q)-1]] <= x {
			q = q[:len(q)-1] // pop队尾
		}
		q = append(q, i) // i放入队列

		left := i - k + 1
		if q[0] < left {
			q = q[1:] // 队首离开队列
		}

		if left >= 0 {
			ans[left] = nums[q[0]] // 队头即最大值
		}
	}
	return ans
}
