package _57_treePaths

import "testing"

//
//  1
///   \
//2     3
//  \
//   5
//
//输出: ["1->2->5", "1->3"]
//
func TestBinaryTreePaths(t *testing.T) {
	node := new(TreeNode)
	node.Val = 1
	node.Left = &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 5}}
	node.Right = &TreeNode{Val: 3}
	t.Log(BinaryTreePaths(node))
	t.Log(BinaryTreePaths2(node))
}
