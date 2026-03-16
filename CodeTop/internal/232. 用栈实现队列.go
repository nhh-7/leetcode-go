package internal

type MyQueue struct {
	stIn, stOut []int
}

func Constructor1() MyQueue {
	return MyQueue{
		stIn:  []int{},
		stOut: []int{},
	}
}

func (mq *MyQueue) Push(x int) {
	mq.stIn = append(mq.stIn, x)
}

func (mq *MyQueue) Pop() int {
	if len(mq.stOut) == 0 {
		for len(mq.stIn) > 0 {
			mq.stOut = append(mq.stOut, mq.stIn[len(mq.stIn)-1])
			mq.stIn = mq.stIn[:len(mq.stIn)-1]
		}
	}
	v := mq.stOut[len(mq.stOut)-1]
	mq.stOut = mq.stOut[:len(mq.stOut)-1]
	return v
}

func (mq *MyQueue) Peek() int {
	if len(mq.stOut) == 0 {
		for len(mq.stIn) > 0 {
			mq.stOut = append(mq.stOut, mq.stIn[len(mq.stIn)-1])
			mq.stIn = mq.stIn[:len(mq.stIn)-1]
		}
	}
	return mq.stOut[len(mq.stOut)-1]
}

func (mq *MyQueue) Empty() bool {
	return len(mq.stIn) == 0 && len(mq.stOut) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
