package _37_averageOfLevels

import "testing"

func Test(t *testing.T) {
	node := new(TreeNode)
	node.Val = 3
	node.Left = &TreeNode{Val: 9, Left: &TreeNode{Val: 15}}
	node.Right = &TreeNode{Val: 20, Right: &TreeNode{Val: 7}}
	t.Log(AverageOfLevels(node))
}
