package hot

import (
	"slices"
	"sort"
)

func minOperations(nums []int, queries []int) []int64 {
	slices.Sort(nums)
	s := make([]int, len(nums)+1)
	for i, x := range nums {
		s[i+1] = s[i] + x
	}

	ans := make([]int64, len(queries))
	for i, q := range queries {
		j := sort.SearchInts(nums, q)
		left := q*j - s[j]
		right := s[len(nums)] - s[j] - q*(len(nums)-j)
		ans[i] = int64(left + right)
	}
	return ans
}
