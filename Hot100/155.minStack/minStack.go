package _55_minStack

//
//设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。
//
//push(x) -- 将元素 x 推入栈中。
//pop() -- 删除栈顶的元素。
//top() -- 获取栈顶元素。
//getMin() -- 检索栈中的最小元素。
//示例:
//
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.getMin();   --> 返回 -2.

//自己写的：
//type MinStack struct {
//	Value []int
//	Min   int
//}
//
///** initialize your data structure here. */
//func Constructor() MinStack {
//	return MinStack{Value: []int{}, Min: 0}
//}
//
//func (this *MinStack) Push(x int) {
//	if len(this.Value) == 0 {
//		this.Min = x
//	} else {
//		if this.Min > x {
//			this.Min = x
//		}
//	}
//
//	this.Value = append(this.Value, x)
//}
//
//func (this *MinStack) Pop() {
//	newList := this.Value
//	this.Value = newList[:len(newList)-1]
//	if len(this.Value) == 0 {
//		return
//	}
//	min := this.Value[0]
//	for _, v := range this.Value {
//		if v < min {
//			min = v
//		}
//	}
//	this.Min = min
//}
//
//func (this *MinStack) Top() int {
//	top := this.Value[len(this.Value)-1]
//	return top
//}
//
//func (this *MinStack) GetMin() int {
//	return this.Min
//}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

//别人写的：
type MinStack struct {
	data   []int
	min    []int
	length int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	minstack := MinStack{
		data:   nil,
		length: 0,
		min:    nil,
	}
	return minstack
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	this.length += 1
	if this.length == 1 || x <= this.min[len(this.min)-1] {
		this.min = append(this.min, x)
	}
}

func (this *MinStack) Pop() {
	this.length -= 1
	if this.data[this.length] == this.min[len(this.min)-1] {
		this.min = this.min[:len(this.min)-1]
	}
	this.data = this.data[:this.length]
}

func (this *MinStack) Top() int {
	return this.data[this.length-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}
