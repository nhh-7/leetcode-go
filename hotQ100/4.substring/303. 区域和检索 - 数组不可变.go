package hot

type NumArray struct {
	s []int
}

func Constructor(nums []int) NumArray {
	na := NumArray{s: make([]int, len(nums)+1)}
	for i, x := range nums {
		na.s[i+1] = na.s[i] + x
	}
	return na
}

func (na *NumArray) SumRange(left int, right int) int {
	return na.s[right+1] - na.s[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
