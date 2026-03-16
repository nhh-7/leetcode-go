package internal

func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}
	ans := 0
	for num := range m {
		if m[num-1] {
			continue
		}

		left := num
		right := left
		for m[right] {
			right++
		}
		ans = max(ans, right-left)
	}
	return ans
}
