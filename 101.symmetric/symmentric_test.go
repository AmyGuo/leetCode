package _01_symmetric

import (
	"fmt"
	"testing"
)

//    1
//   / \
//  2   2
// / \ / \
//3  4 4  3

//   1
//  / \
// 2   2
// /\   \
//   3    3
func TestIsSymmetric(t *testing.T) {
	tree := new(TreeNode)
	tree.Val = 1
	tree.Left = &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}
	tree.Right = &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}
	fmt.Println(IsSymmetric2(tree))

	tree = new(TreeNode)
	tree.Val = 1
	tree.Left = &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}
	tree.Right = &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}
	fmt.Println(IsSymmetric(tree))
}
