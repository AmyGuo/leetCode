package _72_leafSimilar

//请考虑一颗二叉树上所有的叶子，这些叶子的值按从左到右的顺序排列形成一个 叶值序列 。
//
//             3
//           /   \
//          5      1
//         / \     /\
//        6   2   9  8
//           / \
//          7   4
//
//举个例子，如上图所示，给定一颗叶值序列为 (6, 7, 4, 9, 8) 的树。
//
//如果有两颗二叉树的叶值序列是相同，那么我们就认为它们是 叶相似 的。
//
//如果给定的两个头结点分别为 root1 和 root2 的树是叶相似的，则返回 true；否则返回 false 。
//
//
//
//提示：
//
//给定的两颗树可能会有 1 到 200 个结点。
//给定的两颗树上的值介于 0 到 200 之间。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LeafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	arr1 := make([]int, 0)
	arr2 := make([]int, 0)
	dfs(root1, &arr1)
	dfs(root2, &arr2)
	if len(arr1) != len(arr2) {
		return false
	}
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func dfs(node *TreeNode, arr *[]int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		*arr = append(*arr, node.Val)
	}
	dfs(node.Left, arr)
	dfs(node.Right, arr)
}

//相同思路，不同写法
func leafSimilar2(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil || root2 == nil {
		return root2 == root1
	}
	r1 := FindTreeLeaf(root1)
	r2 := FindTreeLeaf(root2)
	if len(r1) != len(r2) {
		return false
	}
	for i, _ := range r1 {
		if r1[i] != r2[i] {
			return false
		}
	}
	return true
}

func FindTreeLeaf(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return []int{}
	}
	if root.Left == nil && root.Right == nil {
		result = append(result, root.Val)
	}
	result = append(result, FindTreeLeaf(root.Left)...)
	result = append(result, FindTreeLeaf(root.Right)...)
	return result
}
