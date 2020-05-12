package _30_getMiniDiff

import "testing"

//  1
//    \
//     3
//     /
//    2
//
func TestGetMinimumDifference(t *testing.T) {
	node := new(TreeNode)
	//node.Val = 1
	//node.Right = &TreeNode{Val: 3, Left: &TreeNode{Val: 2}}
	node.Val = 5
	node.Left = &TreeNode{Val: 4}
	node.Right = &TreeNode{Val: 7}
	t.Log(GetMinimumDifference(node))
	t.Log(GetMinimumDifference2(node))
}
