package _26_invertTree

import (
	"fmt"
	"testing"
)

//     4
//   /   \
//  2     7
// / \   / \
//1   3 6   9
//输出：
//
//     4
//   /   \
//  7     2
// / \   / \
//9   6 3   1

func TestInvertTree(t *testing.T) {
	tree := new(TreeNode)
	tree.Val = 4
	tree.Left = &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	tree.Right = &TreeNode{Val: 7, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 9}}

	PreOrder(tree)
	fmt.Print("-------\n")
	iTree := InvertTree(tree)
	PreOrder(iTree)
}

func PreOrder(node *TreeNode) {
	if node == nil {
		return
	}
	Print(node)
	PreOrder(node.Left)
	PreOrder(node.Right)
}

func Print(node *TreeNode) {
	fmt.Println(node.Val)
}
