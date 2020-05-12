package _83_minDiffInBST

import "testing"

//     4
//   /   \
//   2      6
//  / \
// 1   3
//
func TestMinDiffInBST(t *testing.T) {
	node := new(TreeNode)
	node.Val = 4
	node.Right = &TreeNode{Val: 6}
	node.Left = &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	t.Log(MinDiffInBST(node))
}
