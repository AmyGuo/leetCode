package _90_postOrder

import "testing"

func TestPostOrder(t *testing.T) {
	node := new(Node)
	node.Val = 1
	node.Children = []*Node{&Node{Val: 3, Children: []*Node{&Node{Val: 5}, &Node{Val: 6}}}, &Node{Val: 2}, &Node{Val: 4}}
	t.Log(PostOrder(node))
	t.Log(PostOrder2(node))
}
