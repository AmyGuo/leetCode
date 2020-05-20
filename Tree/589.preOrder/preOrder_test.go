package _89_preOrder

import "testing"

func Test(t *testing.T) {
	node := new(Node)
	node.Val = 1
	node.Children = []*Node{&Node{Val: 3, Children: []*Node{&Node{Val: 5}, &Node{Val: 6}}}, &Node{Val: 2}, &Node{Val: 4}}
	t.Log(PreOrder(node))
	t.Log(PreOrder2(node))
}
