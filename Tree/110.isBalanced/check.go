package _10_isBalanced

//给定一个二叉树，判断它是否是高度平衡的二叉树。
//
//本题中，一棵高度平衡二叉树定义为：
//
//一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。
//
//示例 1:
//
//给定二叉树 [3,9,20,null,null,15,7]
//
//3
/// \
//9  20
//  /  \
//  15   7
//返回 true 。
//
//示例 2:
//
//给定二叉树 [1,2,2,3,3,null,null,4,4]
//
//1
/// \
//2   2
/// \
//3   3
/// \
//4   4
//返回 false 。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftLayer := getLayer(root.Left)
	rightLayer := getLayer(root.Right)

	if leftLayer-rightLayer > 1 || leftLayer-rightLayer < -1 {
		return false
	} else {
		return IsBalanced(root.Left) && IsBalanced(root.Right)
	}
}

func getLayer(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := getLayer(node.Left)
	right := getLayer(node.Right)

	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

func IsBalanced2(root *TreeNode) bool {
	return getDepth(root) != -1
}

func getDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left, right := getDepth(node.Left), getDepth(node.Right)
	if left == -1 || right == -1 {
		return -1
	}
	var depth, max int
	if left > right {
		depth = left - right
		max = left
	} else {
		depth = right - left
		max = right
	}
	if depth > 1 {
		return -1
	}
	return max + 1
}
