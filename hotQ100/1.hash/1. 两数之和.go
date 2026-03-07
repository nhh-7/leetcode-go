package hot

func twoSum1(nums []int, target int) []int {
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

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		if idx, ok := m[target-num]; ok {
			return []int{idx, i}
		}
		m[num] = i
	}
	return []int{}
}
