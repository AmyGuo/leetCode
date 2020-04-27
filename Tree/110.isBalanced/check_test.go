package _10_isBalanced

import (
	"testing"
)

//3
/// \
//9  20
//  /  \
//  15   7

//1
/// \
//2   2
/// \
//3   3
/// \
//4   4

//[1,2,2,3,null,null,3,4,null,null,4]

func TestIsBalanced(t *testing.T) {
	node := new(TreeNode)
	node.Val = 3
	node.Left = &TreeNode{Val: 9}
	node.Right = &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}
	t.Log(IsBalanced(node))
	t.Log(IsBalanced2(node))

	node2 := new(TreeNode)
	node2.Val = 1
	node2.Left = &TreeNode{Val: 2}
	node2.Right = &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}
	t.Log(IsBalanced(node2))
	t.Log(IsBalanced2(node2))

	node3 := new(TreeNode)
	node3.Val = 1
	node3.Left = &TreeNode{Val: 2, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}}
	node3.Right = &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}}}
	t.Log(IsBalanced(node3))
	t.Log(IsBalanced2(node3))
}
