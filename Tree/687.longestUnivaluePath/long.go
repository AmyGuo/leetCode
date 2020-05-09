package _87_longestUnivaluePath

//最长同值路径：
//给定一个二叉树，找到最长的路径，这个路径中的每个节点具有相同值。 这条路径可以经过也可以不经过根节点。
//
//注意：两个节点之间的路径长度由它们之间的边数表示。
//
//示例 1:
//
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
//注意: 给定的二叉树不超过10000个结点。 树的高度不超过1000。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LongestUnivaluePath(root *TreeNode) int {
	ret := 0
	dfs(root, &ret)
	return ret
}

func dfs(node *TreeNode, num *int) int {
	if node == nil {
		return 0
	}
	left := dfs(node.Left, num)
	right := dfs(node.Right, num)

	arrowLeft, arrowRight := 0, 0
	if node.Left != nil && node.Left.Val == node.Val {
		arrowLeft += left + 1
	}
	if node.Right != nil && node.Right.Val == node.Val {
		arrowRight += right + 1
	}

	if *num < arrowRight+arrowLeft {
		*num = arrowRight + arrowLeft
	}

	if arrowRight > arrowLeft {
		return arrowRight
	} else {
		return arrowLeft
	}

}
