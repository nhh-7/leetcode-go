package greedy

import (
	"slices"
)

func reconstructQueue(people [][]int) [][]int {
	/*
		返回 负数：表示 $a$ 应该排在 $b$ 前面。返回 正数：表示 $b$ 应该排在 $a$ 前面。返回 0：表示相等。
	*/
	slices.SortFunc(people, func(a, b []int) int {
		if a[0] == b[0] {
			return a[1] - b[1]
		}
		return b[0] - a[0]
	})
	ans := make([][]int, 0, len(people))
	for _, p := range people {
		idx := p[1]
		ans = append(ans, nil) // 在 result 的末尾强行挤出一个位置，为后续的 copy（向后移动元素）腾出空间。
		copy(ans[idx+1:], ans[idx:])
		ans[idx] = p
	}
	return ans
}
