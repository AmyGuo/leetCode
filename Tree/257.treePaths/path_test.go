package _57_treePaths

import "testing"

func TestBinaryTreePaths(t *testing.T) {
	node := new(TreeNode)
	node.Val = 3
	node.Left = &TreeNode{Val: 9}
	node.Right = &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}
	t.Log(BinaryTreePaths(node))
	t.Log(BinaryTreePaths2(node))
}
