package _87_longestUnivaluePath

import (
	"testing"
)

//输入:
//
//    5
//   / \
//  4   5
// / \   \
// 1   1   5
//输出:
//2

//示例 2:
//输入:
//   1
//  / \
//  4   5
// / \   \
// 4   4   5
//输出:
//2

func TestLongestUnivaluePath(t *testing.T) {
	node1 := new(TreeNode)
	node1.Val = 5
	node1.Left = &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 1}}
	node1.Left = &TreeNode{Val: 5, Right: &TreeNode{Val: 5}}
	t.Log(LongestUnivaluePath(node1))

	node2 := new(TreeNode)
	node2.Val = 1
	node2.Left = &TreeNode{Val: 4, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 4}}
	node2.Right = &TreeNode{Val: 5, Right: &TreeNode{Val: 5}}
	t.Log(LongestUnivaluePath(node2))
}
