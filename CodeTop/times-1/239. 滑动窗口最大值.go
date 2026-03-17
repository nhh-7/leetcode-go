package times1

func maxSlidingWindow(nums []int, k int) []int {
	queue := []int{}

	ans := make([]int, len(nums)-k+1)
	for i, v := range nums {
		for len(queue) > 0 && queue[len(queue)-1] < v {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, v)

		idx := i - k + 1
		if idx >= 0 {
			ans[idx] = queue[0]
			if nums[idx] == queue[0] {
				queue = queue[1:]
			}
		}
	}
	return ans
}
