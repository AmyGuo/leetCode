package _59_MaxDepth

import (
	"testing"
)

//					1
//				 /  |  \
//				3    2   4
//			  /	 \
//			 5	  6
//
func TestMaxDepth(t *testing.T) {
	node := new(Node)
	node.Val = 1
	node.Children = []*Node{&Node{Val: 3, Children: []*Node{&Node{Val: 5}, &Node{Val: 6}}}, &Node{Val: 2}, &Node{Val: 4}}
	t.Log(MaxDepth(node))
	t.Log(MaxDepth2(node))
}
