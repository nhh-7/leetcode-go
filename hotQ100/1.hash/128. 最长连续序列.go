package hot

func longestConsecutive1(nums []int) int {
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

func longestConsecutive(nums []int) int {
	m := map[int]bool{}
	for _, num := range nums {
		m[num] = true
	}

	ans := 0
	// 为了防止原始数组中的重复项太多，选择遍历map而不是原数组
	for num := range m {
		if m[num-1] {
			continue
		}

		tmp := num
		for m[tmp+1] {
			tmp++
		}
		ans = max(tmp-num+1, ans)
	}
	return ans
}
