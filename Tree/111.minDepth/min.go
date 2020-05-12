package _11_minDepth

//111. 二叉树的最小深度
//给定一个二叉树，找出其最小深度。
//
//最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
//
//说明: 叶子节点是指没有子节点的节点。
//
//示例:
//
//给定二叉树 [3,9,20,null,null,15,7],
//
//  3
// / \
// 9  20
//   /  \
//   15   7
//返回它的最小深度  2.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	nodes := []*TreeNode{root}
	min := 0
	for len(nodes) > 0 {
		temp := []*TreeNode{}
		min++
		for _, v := range nodes {
			if v.Left == nil && v.Right == nil {
				return min
			}
			if v.Left != nil {
				temp = append(temp, v.Left)
			}
			if v.Right != nil {
				temp = append(temp, v.Right)
			}
		}
		nodes = temp
	}
	return min
}

//解决办法：
//1.遍历二叉树，记录其深度，到叶子节点时返回深度
//2.比较左右字数的深度，返回最小值
//3.如果没有某个子树时，不比较
func MinDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		return 1
	}

	return minDepthInner(root, 1)
}

func minDepthInner(root *TreeNode, level int) int {
	if root.Left == nil && root.Right == nil {
		return level
	}

	lleft := 0
	lright := 0
	if root.Left != nil {
		lleft = minDepthInner(root.Left, level+1)
	}
	if root.Right != nil {
		lright = minDepthInner(root.Right, level+1)
	}

	if lright*lleft == 0 {
		return lright + lleft
	} else {
		if lleft < lright {
			return lleft
		} else {
			return lright
		}
	}
}
