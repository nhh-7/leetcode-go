package hot

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		idx, ok := m[target-v]
		if ok {
			return []int{i, idx}
		}
		m[v] = i
	}
	return []int{0, 0}

}
