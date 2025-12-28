package hot

func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	ans := 0
	for _, v := range nums {
		m[v] = true
	}

	for x := range m {
		if m[x-1] == true {
			continue
		}

		y := x
		for m[y+1] {
			y++
		}

		ans = max(ans, y-x+1)
	}
	return ans
}
