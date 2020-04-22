package _04_maxDepth

import (
	"fmt"
	"testing"
)

//3
/// \
//9  20
///  \
//15   7

func TestMaxDepth(t *testing.T) {
	tree := new(TreeNode)
	tree.Val = 3
	tree.Left = &TreeNode{Val: 9, Left: &TreeNode{Val: 15}}
	tree.Right = &TreeNode{Val: 20, Right: &TreeNode{Val: 7}}
	fmt.Println(MaxDepth(tree))
	fmt.Println(MaxDepth2(tree))
}
