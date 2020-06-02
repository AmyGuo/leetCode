package _12_hasPathSum

import (
	"testing"
)

//
//    5
//   / \
//  4    8
// /     / \
// 11    13  4
// /  \       \
// 7    2       1

func TestHasPathSum(t *testing.T) {
	node := new(TreeNode)
	node.Val = 5
	node.Left = &TreeNode{Val: 4, Left: &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}}}
	node.Right = &TreeNode{Val: 8, Left: &TreeNode{Val: 13}, Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 1}}}
	t.Log(HasPathSum(node, 22))
	t.Log(HasPathSum2(node, 22))
}
