package _30_getMiniDiff

import "math"

//二叉搜索树的最小绝对差
//给你一棵所有节点为非负值的二叉搜索树，请你计算树中任意两节点的差的绝对值的最小值。
//
//
//
//示例：
//
//输入：
//
//  1
//    \
//     3
//     /
//    2
//
//输出：
//1
//
//解释：
//最小绝对差为 1，其中 2 和 1 的差的绝对值为 1（或者 2 和 3）。
//
//
//提示：
//
//树中至少有 2 个节点。
//本题与 783 https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/ 相同

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//思路没错，中序遍历是升序列，但是新开辟了空间，可优化掉
func GetMinimumDifference(root *TreeNode) int {
	nums := make([]int, 0)
	dfs(root, &nums)
	min := math.MaxInt64
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] < min {
			min = nums[i] - nums[i-1]
		}
	}

	return min
}

func dfs(node *TreeNode, nums *[]int) {
	if node == nil {
		return
	}
	dfs(node.Left, nums)
	*nums = append(*nums, node.Val)
	dfs(node.Right, nums)
}

func GetMinimumDifference2(root *TreeNode) int {
	pre := -1
	min := math.MaxInt64
	midSort(root, &pre, &min)
	return min
}

func midSort(root *TreeNode, pre, min *int) {
	if root == nil {
		return
	}
	midSort(root.Left, pre, min)
	if *pre != -1 {
		if root.Val-*pre < *min {
			*min = root.Val - *pre
		}
	}
	*pre = root.Val
	midSort(root.Right, pre, min)
}
