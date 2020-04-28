package _01_findMod

import "testing"

func TestFindMode(t *testing.T) {
	node := new(TreeNode)
	node.Val = 1
	node.Right = &TreeNode{Val: 2, Left: &TreeNode{Val: 2}}
	t.Log(FindMode(node))
}
