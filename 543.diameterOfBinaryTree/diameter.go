package _43_diameterOfBinaryTree

//给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过根结点。
//
//示例 :
//给定二叉树
//
//1
/// \
//2   3
/// \
//4   5
//返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。
//
//注意：两结点之间的路径长度是以它们之间边的数目表示。
//

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//分析：这道题看似求直径，其实是求深度的变形。深度 = max(左子数深度，右子树深度) +1

//class Solution:
//
//def __init__(self):
//self.max = 0
//
//def diameterOfBinaryTree(self, root: TreeNode) -> int:
//self.depth(root)
//
//return self.max
//
//def depth(self, root):
//if not root:
//return 0
//l = self.depth(root.left)
//r = self.depth(root.right)
//'''每个结点都要去判断左子树+右子树的高度是否大于self.max，更新最大值'''
//self.max = max(self.max, l+r)
//
//# 返回的是高度
//return max(l, r) + 1

func DiameterOfBinaryTree(root *TreeNode) int {
	max := 0
	depth(root, &max)
	return max
}

func depth(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}

	leftDepth := depth(root.Left, max)
	rightDepth := depth(root.Right, max)
	if leftDepth+rightDepth > *max {
		*max = leftDepth + rightDepth
	}

	if rightDepth > leftDepth {
		return rightDepth + 1
	} else {
		return leftDepth + 1
	}
}
