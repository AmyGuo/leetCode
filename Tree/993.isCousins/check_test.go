package _93_isCousins

import "testing"

////        1
//         / \
//        2    3
//         \    \
//          4    5
func TestIsCousins(t *testing.T) {
	node := new(TreeNode)
	node.Val = 1
	node.Left = &TreeNode{Val: 2, Right: &TreeNode{Val: 4}}
	node.Right = &TreeNode{Val: 3, Right: &TreeNode{Val: 5}}
	t.Log(IsCousins(node, 5, 4))
	t.Log(IsCousins2(node, 5, 4))

	node1 := new(TreeNode)
	node1.Val = 1
	node1.Left = &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}
	node1.Right = &TreeNode{Val: 3}
	t.Log(IsCousins(node1, 4, 3))
	t.Log(IsCousins2(node1, 4, 3))
}
