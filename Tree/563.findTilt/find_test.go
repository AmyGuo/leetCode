package _63_findTilt

import "testing"

func TestFindTile(t *testing.T) {
	node := new(TreeNode)
	node.Val = 1
	node.Left = &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 5}}
	node.Right = &TreeNode{Val: 3}
	t.Log(FindTile(node))
}
