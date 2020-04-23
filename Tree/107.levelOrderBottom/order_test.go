package _07_levelOrderBottom

import "testing"

//3
/// \
//9  20
//  /  \
//  15   7
//[3,9,20,null,null,15,7],
func TestLevelOrderBottom(t *testing.T) {
	node := new(TreeNode)
	node.Val = 3
	node.Left = &TreeNode{Val: 9}
	node.Right = &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}
	t.Log(LevelOrderBottom(node))
	t.Log(LevelOrderBottom2(node))
}
