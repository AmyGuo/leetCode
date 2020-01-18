package _38_convertBST

import (
	"fmt"
	"testing"
)

//输入: 二叉搜索树:
//5
///   \
//2     13
//
//输出: 转换为累加树:
//18
///   \
//20     13

func TestConvertBST(t *testing.T) {
	tree := new(TreeNode)
	tree.Val = 5
	tree.Left = &TreeNode{Val: 2}
	tree.Right = &TreeNode{Val: 13}
	PreOrder(tree)
	nTree := ConvertBST(tree)
	PreOrder(nTree)
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
