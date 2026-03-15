package internal

func maxSlidingWindow(nums []int, k int) []int {
	queue := []int{}

	ans := make([]int, len(nums)-k+1)
	for i, v := range nums {
		for len(queue) > 0 && queue[len(queue)-1] < v {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, v)
		if i-k+1 >= 0 {
			ans[i-k+1] = queue[0]
			if queue[0] == nums[i-k+1] {
				queue = queue[1:]
			}
		}
	}
	return ans
}
