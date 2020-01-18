package _43_diameterOfBinaryTree

import (
	"fmt"
	"testing"
)

//1
/// \
//2   3
/// \
//4   5
//返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。

func TestDiameterOfBinaryTree(t *testing.T) {
	tree := new(TreeNode)
	tree.Val = 1
	//tree.Left = &TreeNode{Val: 2, Right: &TreeNode{Val: 5}, Left: &TreeNode{Val: 4}}
	//tree.Right = &TreeNode{Val: 3}
	fmt.Println(DiameterOfBinaryTree(tree))
}
