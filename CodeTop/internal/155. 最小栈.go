package internal

import "math"

type pair struct {
	val, mmin int
}

type MinStack struct {
	stack []pair
}

func Constructor() MinStack {
	return MinStack{stack: []pair{{0, math.MaxInt}}}
}

func (ms *MinStack) Push(val int) {
	curMin := min(val, ms.GetMin())
	p := pair{val: val, mmin: curMin}
	ms.stack = append(ms.stack, p)
}

func (ms *MinStack) Pop() {
	ms.stack = ms.stack[:len(ms.stack)-1]
}

func (ms *MinStack) Top() int {
	return ms.stack[len(ms.stack)-1].val
}

func (ms *MinStack) GetMin() int {
	return ms.stack[len(ms.stack)-1].mmin
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
