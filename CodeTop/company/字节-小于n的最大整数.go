package main

import (
	"fmt"
	"sort"
	"strconv"
)

func maxLess(n int, A []int) int {
	sort.Ints(A)
	s := strconv.Itoa(n)
	m := len(s)
	res := make([]int, m)

	for i := 0; i < m; i++ {
		cur := int(s[i] - '0')

		// 找 <= cur 的最大数字
		j := sort.SearchInts(A, cur+1) - 1
		if j >= 0 {
			res[i] = A[j]
			if A[j] < cur {
				for k := i + 1; k < m; k++ {
					res[k] = A[len(A)-1]
				}
				return toInt(res)
			}
		} else {
			// 回退
			for k := i - 1; k >= 0; k-- {
				j := sort.SearchInts(A, res[k]) - 1
				if j >= 0 {
					res[k] = A[j]
					for t := k + 1; t < m; t++ {
						res[t] = A[len(A)-1]
					}
					return toInt(res)
				}
			}
			// 位数减少
			ans := 0
			for i := 0; i < m-1; i++ {
				ans = ans*10 + A[len(A)-1]
			}
			return ans
		}
	}
	return toInt(res)
}

func toInt(a []int) int {
	x := 0
	for _, v := range a {
		x = x*10 + v
	}
	return x
}

func main() {
	fmt.Println(maxLess(23121, []int{2, 4, 9})) // 22999
}
