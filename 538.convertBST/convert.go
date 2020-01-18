package _38_convertBST

//
//给定一个二叉搜索树（Binary Search Tree），把它转换成为累加树（Greater Tree)，使得每个节点的值是原来的节点值加上所有大于它的节点值之和。
//
//例如：
//
//输入: 二叉搜索树:
//5
///   \
//2     13
//
//输出: 转换为累加树:
//18
///   \
//20     13

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//分析：每个节点是原来节点加上所有大于它的节点值之和，因为二叉树的右节点总是大于根节点，所以就是要计算每个节点的右子树节点之和
//即变相的中序遍历 (右-根-左)
func ConvertBST(root *TreeNode) *TreeNode {
	n := 0
	dfs(root, &n)
	return root
}

func dfs(root *TreeNode, num *int) {
	if root == nil {
		return
	}
	dfs(root.Right, num)
	root.Val += *num
	*num = root.Val
	dfs(root.Left, num)
}
