package _26_invertTree

//翻转一棵二叉树。
//
//示例：
//
//输入：
//
//     4
//   /   \
//  2     7
// / \   / \
//1   3 6   9
//输出：
//
//     4
//   /   \
//  7     2
// / \   / \
//9   6 3   1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := InvertTree(root.Left)
	right := InvertTree(root.Right)
	root.Right = left
	root.Left = right
	return root
}
