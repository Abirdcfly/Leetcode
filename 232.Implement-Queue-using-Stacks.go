type Stack struct {
	d []int
	l int
}

/** initialize your data structure here. */
func ConstructorStack() Stack {
	return Stack{
		d: make([]int, 0),
		l: 0,
	}
}

func (this *Stack) Push(x int) {
	this.d = append(this.d, x)
	this.l++
}

func (this *Stack) Pop() int {
	x := this.d[len(this.d)-1]
	this.d = this.d[:len(this.d)-1]
	this.l--
	return x
}

func (this *Stack) Top() int {
	return this.d[len(this.d)-1]
}

func (this *Stack) Empty() bool {
	return this.l == 0
}

type MyQueue struct {
	instack  *Stack
	outstack *Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	in := ConstructorStack()
	out := ConstructorStack()
	return MyQueue{
		instack:  &in,
		outstack: &out,
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.instack.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.outstack.Empty() {
		for !this.instack.Empty() {
			this.outstack.Push(this.instack.Pop())
		}
	}
	return this.outstack.Pop()
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.outstack.Empty() {
		for !this.instack.Empty() {
			this.outstack.Push(this.instack.Pop())
		}
	}
	return this.outstack.Top()
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.instack.Empty() && this.outstack.Empty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
