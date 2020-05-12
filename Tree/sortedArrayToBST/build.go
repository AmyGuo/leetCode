package sortedArrayToBST

//面试题 04.02. 最小高度树
//给定一个有序整数数组，元素各不相同且按升序排列，编写一个算法，创建一棵高度最小的二叉搜索树。
//
//示例:
//给定有序数组: [-10,-3,0,5,9],
//
//一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：
//
//		      0
//			/   \
//			-3   9
//         /   /
//        -10  5

//同108题
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	left := nums[:mid]
	right := nums[mid+1:]
	node := &TreeNode{Val: nums[mid], Left: sortedArrayToBST(left), Right: sortedArrayToBST(right)}
	return node
}
