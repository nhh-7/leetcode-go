package times1

type MyQueue struct {
	stkOut, stkIn []int
}

func Constructor() MyQueue {
	return MyQueue{
		stkOut: []int{},
		stkIn:  []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.stkIn = append(this.stkIn, x)
}

func (this *MyQueue) Pop() int {
	if len(this.stkOut) == 0 {
		for len(this.stkIn) > 0 {
			this.stkOut = append(this.stkOut, this.stkIn[len(this.stkIn)-1])
			this.stkIn = this.stkIn[:len(this.stkIn)-1]
		}
	}
	v := this.stkOut[len(this.stkOut)-1]
	this.stkOut = this.stkOut[:len(this.stkOut)-1]
	return v
}

func (this *MyQueue) Peek() int {
	if len(this.stkOut) == 0 {
		for len(this.stkIn) > 0 {
			this.stkOut = append(this.stkOut, this.stkIn[len(this.stkIn)-1])
			this.stkIn = this.stkIn[:len(this.stkIn)-1]
		}
	}
	return this.stkOut[len(this.stkOut)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.stkIn) == 0 && len(this.stkOut) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
