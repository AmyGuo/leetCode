package _17_mergeTrees

import (
	"fmt"
	"testing"
)

//输入:
//Tree 1                     Tree 2
//1                         2
/// \                       / \
//3   2                     1   3
///                           \   \
//5                             4   7
//输出:
//合并后的树:
//3
/// \
//4   5
/// \   \
//5   4   7
func TestMergeTrees(t *testing.T) {
	t1 := new(TreeNode)
	t1.Val = 1
	t1.Left = &TreeNode{Val: 3, Left: &TreeNode{Val: 5}}
	t1.Right = &TreeNode{Val: 2}

	t2 := new(TreeNode)
	t2.Val = 2
	t2.Left = &TreeNode{Val: 1, Right: &TreeNode{Val: 4}}
	t2.Right = &TreeNode{Val: 3, Right: &TreeNode{Val: 7}}
	//PreOrder(t1)
	//fmt.Println("-------")
	//PreOrder(t2)
	//fmt.Println("-------")

	t3 := MergeTrees(t1, t2)
	PreOrder(t3)
}

func PreOrder(node *TreeNode) {
	if node == nil {
		return
	}
	Print(node)
	PreOrder(node.Left)
	PreOrder(node.Right)
}

func Print(node *TreeNode) {
	fmt.Println(node.Val)
}
