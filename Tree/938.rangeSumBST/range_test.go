package _38_rangeSumBST

import "testing"

func TestRangeSumBST(t *testing.T) {
	node := new(TreeNode)
	node.Val = 10
	node.Left = &TreeNode{Val: 5, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 7}}
	node.Right = &TreeNode{Val: 15, Right: &TreeNode{Val: 18}}
	left := 7
	right := 15

	t.Log(RangeSumBST(node, left, right))
	t.Log(RangeSumBST2(node, left, right))
}
