package _63_findTilt

//给定一个二叉树，计算整个树的坡度。
//
//一个树的节点的坡度定义即为，该节点左子树的结点之和和右子树结点之和的差的绝对值。空结点的的坡度是0。
//
//整个树的坡度就是其所有节点的坡度之和。
//
//示例:
//
//输入:
//   1
///   \
//2     3
//输出: 1
//解释:
//结点的坡度 2 : 0
//结点的坡度 3 : 0
//结点的坡度 1 : |2-3| = 1
//树的坡度 : 0 + 0 + 1 = 1
//注意:
//
//任何子树的结点的和不会超过32位整数的范围。
//坡度的值不会超过32位整数的范围。
//

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func FindTile(root *TreeNode) int {
	ret := 0
	traverse(root, &ret)
	return ret
}

func traverse(node *TreeNode, ret *int) int {
	if node == nil {
		return 0
	}

	left := traverse(node.Left, ret)
	right := traverse(node.Right, ret)
	*ret += abs(left, right)
	return left + right + node.Val
}

func abs(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
