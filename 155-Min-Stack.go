//方法有2种，一种是这种最小值入栈还有是用tuple(入栈，未加入此元素时的最小值)
import "math"

type MinStack struct {
	d   []int
	min int
	l   int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		d:   make([]int, 0),
		min: math.MaxInt32,
		l:   0,
	}
}

func (this *MinStack) Push(x int) {
	if x <= this.min {
		this.d = append(this.d, this.min)
		this.min = x
	}
	this.d = append(this.d, x)
	this.l++
}

func (this *MinStack) Pop() {
	if this.d[len(this.d)-1] == this.min {
		this.d = this.d[:len(this.d)-1]
		this.min = this.d[len(this.d)-1]
	}
	this.d = this.d[:len(this.d)-1]
	this.l--
}

func (this *MinStack) Top() int {
	return this.d[len(this.d)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
