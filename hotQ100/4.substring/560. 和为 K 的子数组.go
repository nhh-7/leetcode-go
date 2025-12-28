package hot

func subarraySum(nums []int, k int) int {
	ans := 0
	m := make(map[int]int, len(nums)+1)
	m[0] = 1
	s := 0
	for _, x := range nums {
		s += x
		ans += m[s-k]
		m[s]++
	}
	return ans
}
