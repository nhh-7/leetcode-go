package stack

// 单调栈内递增，寻找下一个更小的
// 以[4,4,4,3]为例，当遍历到下标0，以4为高度，那么需要向右扩展，找到第一个比4小的数，以他来作为矩形的右边界
// 而左边界就是该下标在单调栈中的上一个元素
func largestRectangleArea(heights []int) int {
	stack := []int{}
	stack = append(stack, -1)    // 哨兵结点，用于处理以第一个元素为高
	heights = append(heights, 0) // 哨兵
	ans := 0

	for i, t := range heights {
		// 一个递增的序列会先进入栈，然后从栈顶开始逐个作为矩形高进行枚举计算
		// [4，5，6，1]  i遍历到3时，开始计算矩形面积，在这个过程中i（右边界）不变，而左边界和高度在变化（内层for循环）
		for len(stack) > 1 && t < heights[stack[len(stack)-1]] {
			// top是矩形高度
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 左边界, 当前遍历到的i是右边界
			l := stack[len(stack)-1]
			ans = max(ans, (i-l-1)*heights[top])
		}
		stack = append(stack, i)
	}
	return ans
}
