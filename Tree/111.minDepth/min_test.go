package _11_minDepth

import "testing"

// 3
/// \
//9  20
//  /  \
//  15   7
func TestMinDepth(t *testing.T) {
	node := new(TreeNode)
	node.Val = 3
	node.Left = &TreeNode{Val: 9}
	node.Right = &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}

	node2 := new(TreeNode)
	node2.Val = 1
	node2.Left = &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}
	node2.Right = &TreeNode{Val: 3}
	t.Log(MinDepth(node))
	t.Log(MinDepth2(node))
	t.Log("------------------------------")
	t.Log(MinDepth(node2))
	t.Log(MinDepth2(node2))
}
