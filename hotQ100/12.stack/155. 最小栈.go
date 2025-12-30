package stack

// 使用两个栈，一个栈的栈顶用于记录当前栈中的最小元素，另一个为正常的栈
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

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
