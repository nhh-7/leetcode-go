package skills

func sortColors(nums []int) {
	p0, curr := 0, 0
	p2 := len(nums) - 1
	for curr <= p2 {
		if nums[curr] == 0 {
			nums[p0], nums[curr] = nums[curr], nums[p0]
			p0++
			curr++
		} else if nums[curr] == 2 {
			nums[curr], nums[p2] = nums[p2], nums[curr]
			p2--
			// 这里 curr 不动，因为从 p2 换回来的数可能是 0 或 1，需要下一轮检查
		} else {
			curr++
		}
	}
}
