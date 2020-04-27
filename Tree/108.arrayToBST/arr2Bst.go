package _08_arrayToBST

//将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。
//
//本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
//
//示例:
//
//给定有序数组: [-10,-3,0,5,9],
//
//一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：
//
//	    0
//	   / \
//   -3   9
//   /    /
// -10   5

//注意：二叉树的一些知识点
//中序遍历不能唯一确定一棵二叉搜索树。
//先序和后序遍历不能唯一确定一棵二叉搜索树。
//先序/后序遍历和中序遍历的关系：
//inorder = sorted(postorder) = sorted(preorder)，
//中序+后序、中序+先序可以唯一确定一棵二叉树。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2

	left := nums[:mid]
	right := nums[mid+1:]
	node := &TreeNode{Val: nums[mid], Left: SortedArrayToBST(left), Right: SortedArrayToBST(right)}
	return node
}
