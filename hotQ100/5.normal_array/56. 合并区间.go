package hot

import (
	"slices"
)

func merge(intervals [][]int) [][]int {
	ans := [][]int{}
	slices.SortFunc(intervals, func(lhs, rhs []int) int {
		if lhs[0] == rhs[0] {
			return lhs[1] - rhs[1]
		}
		return lhs[0] - rhs[0]
	})

	left, right := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if right < intervals[i][0] {
			ans = append(ans, []int{left, right})
			left = intervals[i][0]
			right = intervals[i][1]
		} else {
			right = max(right, intervals[i][1])
		}
	}
	ans = append(ans, []int{left, right})
	return ans
}
