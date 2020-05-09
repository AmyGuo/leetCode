package base

import (
	"fmt"
	"testing"
)

func TestCreateTreeNode(t *testing.T) {
	root := TreeNode{Value: 3}
	root.Left = &TreeNode{}
	root.Left.SetValue(0)
	root.Left.Right = CreateTreeNode(2)
	root.Right = &TreeNode{5, nil, nil}
	root.Right.Left = CreateTreeNode(4)

	fmt.Print("\n前序遍历: ")
	root.PreOrder()
	fmt.Print("\n中序遍历: ")
	root.MiddleOrder()
	fmt.Print("\n后序遍历: ")
	root.PostOrder()
	fmt.Print("\n层次遍历: ")
	root.BreadthFirstSearch()
	fmt.Println("\n层数: ", root.Layers())
	fmt.Println("\n层数: ", root.LayersByQueue())

	node := BuildTreeNode([]interface{}{6, 2, 8, 0, 4, 7, 9})
	node.BreadthFirstSearch()
}
