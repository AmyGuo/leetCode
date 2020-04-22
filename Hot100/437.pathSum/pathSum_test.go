package _37_pathSum

import (
	"fmt"
	"testing"
)

//10
///  \
//5   -3
/// \    \
//3   2   11
/// \   \
//3  -2   1
func TestPathSum(t *testing.T) {
	tree := new(TreeNode)
	tree.Val = 10
	tree.Left = &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: -2}}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 1}}}
	tree.Right = &TreeNode{Val: -3, Right: &TreeNode{Val: 11}}
	fmt.Println(PathSum(tree, 7))
}
