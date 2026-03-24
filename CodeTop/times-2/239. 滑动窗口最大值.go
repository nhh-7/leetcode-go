package times2

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	ans := make([]int, n-k+1)
	queue := []int{}
	for right, v := range nums {
		for len(queue) > 0 && queue[len(queue)-1] < v {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, v)
		left := right - k + 1
		if left >= 0 {
			ans[left] = queue[0]
			if queue[0] == nums[left] {
				queue = queue[1:]
			}
		}
	}
	return ans
}
