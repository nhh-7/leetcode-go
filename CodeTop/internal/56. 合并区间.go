package internal

import "slices"

func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a, b []int) int {
		if a[0] == b[0] {
			return a[1] - b[1]
		}
		return a[0] - b[0]
	})

	ans := [][]int{}
	left := intervals[0][0]
	right := intervals[0][1]
	for _, inte := range intervals {
		if inte[0] <= right {
			right = max(right, inte[1])
		} else {
			ans = append(ans, []int{left, right})
			left = inte[0]
			right = inte[1]
		}
	}
	ans = append(ans, []int{left, right})
	return ans
}
