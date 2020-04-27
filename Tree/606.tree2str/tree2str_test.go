package _06_tree2str

import "testing"

func TestTree2str(t *testing.T) {
	node := new(TreeNode)
	node.Val = 1
	node.Right = &TreeNode{Val: 3}
	node.Left = &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}
	t.Log(Tree2str(node))

	node1 := new(TreeNode)
	node1.Val = 1
	node1.Right = &TreeNode{Val: 3}
	node1.Left = &TreeNode{Val: 2, Right: &TreeNode{Val: 4}}
	t.Log(Tree2str(node1))
}
