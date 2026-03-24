package times2

type MyQueue struct {
	stkIn, stkOut []int
}

func Constructor() MyQueue {
	return MyQueue{
		stkIn:  []int{},
		stkOut: []int{},
	}
}

func (mq *MyQueue) Push(x int) {
	mq.stkIn = append(mq.stkIn, x)
}

func (mq *MyQueue) Pop() int {
	if len(mq.stkOut) == 0 {
		for len(mq.stkIn) > 0 {
			mq.stkOut = append(mq.stkOut, mq.stkIn[len(mq.stkIn)-1])
			mq.stkIn = mq.stkIn[:len(mq.stkIn)-1]
		}
	}
	v := mq.stkOut[len(mq.stkOut)-1]
	mq.stkOut = mq.stkOut[:len(mq.stkOut)-1]
	return v
}

func (mq *MyQueue) Peek() int {
	if len(mq.stkOut) == 0 {
		for len(mq.stkIn) > 0 {
			mq.stkOut = append(mq.stkOut, mq.stkIn[len(mq.stkIn)-1])
			mq.stkIn = mq.stkIn[:len(mq.stkIn)-1]
		}
	}
	v := mq.stkOut[len(mq.stkOut)-1]
	return v
}

func (mq *MyQueue) Empty() bool {
	return len(mq.stkIn) == 0 && len(mq.stkOut) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
