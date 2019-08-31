//https://leetcode.com/problems/implement-stack-using-queues/solution/
type MyStack struct {
	d []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		d: make([]int, 0),
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.d = append(this.d, x)
	for i := 0; i < len(this.d); i++ {
		this.d = append(this.d, this.Pop())
	}
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	res := 0
	res, this.d = this.d[len(this.d)-1], this.d[:len(this.d)-1]
	return res
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.d[len(this.d)-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.d) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
