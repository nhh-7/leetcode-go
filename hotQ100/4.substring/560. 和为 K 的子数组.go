package hot

func subarraySum1(nums []int, k int) int {
	s := make([]int, len(nums)+1)
	for i, num := range nums {
		s[i+1] = s[i] + num
	}

	ans := 0
	cnt := make(map[int]int, len(nums)+1)
	for _, sj := range s {
		ans += cnt[sj-k]
		cnt[sj]++
	}
	return ans
}

func subarraySum(nums []int, k int) int {
	ans := 0
	cnt := make(map[int]int, len(nums)+1)
	cnt[0] = 1
	s := 0
	for _, x := range nums {
		s += x
		ans += cnt[s-k]
		cnt[s]++
	}
	return ans
}
