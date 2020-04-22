package _55_minStack

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	//MinStack minStack = new MinStack();
	//minStack.push(-2);
	//minStack.push(0);
	//minStack.push(-3);
	//minStack.getMin();   --> 返回 -3.
	//minStack.pop();
	//minStack.top();      --> 返回 0.
	//minStack.getMin();   --> 返回 -2.
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
