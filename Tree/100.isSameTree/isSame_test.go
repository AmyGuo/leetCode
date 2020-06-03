package _00_isSameTree

import "testing"

//输入:       1         1
//           / \       / \
//          2   3     2   3
//
//[1,2,3],   [1,2,3]
//
//输出: true
//示例 2:
//
//输入:      1          1
//			/           \
//			2             2
//
//[1,2],     [1,null,2]

func TestIsSameTree(t *testing.T) {
	node1 := new(TreeNode)
	node2 := new(TreeNode)
	node1.Val = 1
	node2.Val = 1
	node1.Left = &TreeNode{Val: 2}
	node2.Left = &TreeNode{Val: 2}
	node1.Right = &TreeNode{Val: 3}
	node2.Right = &TreeNode{Val: 3}

	node3 := new(TreeNode)
	node4 := new(TreeNode)
	node3.Val = 1
	node4.Val = 1
	node3.Left = &TreeNode{Val: 2}
	node4.Right = &TreeNode{Val: 2}

	t.Log(IsSameTree(node1, node2))
	t.Log(IsSameTree2(node1, node2))

	t.Log(IsSameTree(node3, node4))
	t.Log(IsSameTree2(node3, node4))
}
