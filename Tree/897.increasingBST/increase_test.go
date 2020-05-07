package _97_increasingBST

import (
	"fmt"
	"testing"
)

func TestIncreasingBST(t *testing.T) {
	node := new(TreeNode)
	node.Val = 5
	node.Left = &TreeNode{Val: 3, Right: &TreeNode{Val: 4}, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}}
	node.Right = &TreeNode{Val: 6, Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 9}}}
	//IncreasingBST(node).PreOrder()
	IncreasingBST2(node).PreOrder()
}

func (node *TreeNode) PreOrder() {
	if node == nil {
		return
	}
	node.Print()
	node.Left.PreOrder()
	node.Right.PreOrder()
}
func (node *TreeNode) Print() {
	fmt.Print(node.Val, " ")
}
