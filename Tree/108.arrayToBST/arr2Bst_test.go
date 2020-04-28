package _08_arrayToBST

import (
	"fmt"
	"testing"
)

func TestSortedArrayToBST(t *testing.T) {
	node := SortedArrayToBST([]int{-10, -3, 0, 5, 9})
	node.PreOrder()
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
