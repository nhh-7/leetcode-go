package stack

import "math"

/*
使用两个栈，一个栈的栈顶用于记录当前栈中的最小元素，另一个为正常的栈
type MinStack struct {
	minStk []int
	stk    []int
}

func Constructor() MinStack {
	return MinStack{
		minStk: []int{},
		stk:    []int{},
	}
}

func (m *MinStack) Push(val int) {
	m.stk = append(m.stk, val)
	if len(m.minStk) == 0 || val <= m.minStk[len(m.minStk)-1] {
		m.minStk = append(m.minStk, val)
	}
}

func (m *MinStack) Pop() {
	if m.Top() == m.minStk[len(m.minStk)-1] {
		m.minStk = m.minStk[:len(m.minStk)-1]
	}
	m.stk = m.stk[:len(m.stk)-1]
}

func (m *MinStack) Top() int {
	return m.stk[len(m.stk)-1]
}

func (m *MinStack) GetMin() int {
	return m.minStk[len(m.minStk)-1]
}
*/

// 使用一个pair保存当前栈中的元素及最小值
// type pair struct {
// 	val, preMin int
// }
// type MinStack []pair

// func Constructor() MinStack {
// 	return MinStack{{0, math.MaxInt}}
// }

// func (m *MinStack) Push(val int) {
// 	curMin := min(val, m.GetMin())
// 	*m = append(*m, pair{val: val, preMin: curMin})
// }

// func (m *MinStack) Pop() {
// 	*m = (*m)[:len(*m)-1]
// }

// func (m *MinStack) Top() int {
// 	return (*m)[len(*m)-1].val
// }

// func (m *MinStack) GetMin() int {
// 	return (*m)[len(*m)-1].preMin
// }

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

type pair struct {
	val, mmin int
}

type MinStack struct {
	stack []pair
}

func Constructor() MinStack {
	return MinStack{
		stack: []pair{{0, math.MaxInt}},
	}
}

func (this *MinStack) Push(val int) {
	curMin := min(val, this.GetMin())
	this.stack = append(this.stack, pair{val: val, mmin: curMin})
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1].val
}

func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack)-1].mmin
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
