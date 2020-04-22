package _01_symmetric

import (
	"errors"
)

//给定一个二叉树，检查它是否是镜像对称的。
//
//例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
//
//    1
//   / \
//  2   2
// / \ / \
//3  4 4  3
//但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
//
//   1
//  / \
// 2   2
// /\   \
//   3    3

//解析：
//
//它们的两个根结点具有相同的值。
//每个树的右子树都与另一个树的左子树镜像对称

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归法：
func IsSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}

	if t1 == nil || t2 == nil {
		return false
	}

	return t1.Val == t2.Val && isMirror(t1.Right, t2.Left) && isMirror(t1.Left, t2.Right)
}

//复杂度分析
//
//时间复杂度：O(n)，因为我们遍历整个输入树一次，所以总的运行时间为 O(n)，其中 n 是树中结点的总数。
//空间复杂度：递归调用的次数受树的高度限制。在最糟糕情况下，树是线性的，其高度为 O(n)。因此，在最糟糕的情况下，由栈上的递归调用造成的空间复杂度为 O(n)。

//迭代法：

type NodeList []*TreeNode

func (n *NodeList) IsEmpty() bool {
	return len(*n) == 0
}

func (n *NodeList) Pop() (*TreeNode, error) {
	if n.IsEmpty() {
		return &TreeNode{}, errors.New("empty")
	}
	temp := *n
	value := temp[len(temp)-1]
	*n = temp[:len(temp)-1]
	return value, nil
}

func (n *NodeList) Push(node *TreeNode) {
	*n = append(*n, node)
}

func IsSymmetric2(root *TreeNode) bool {
	list := make(NodeList, 0)
	list.Push(root)
	list.Push(root)
	for !list.IsEmpty() {
		t1, _ := list.Pop()
		t2, _ := list.Pop()
		if t1 == nil && t2 == nil {
			continue
		}
		if t1 == nil || t2 == nil {
			return false
		}
		if t2.Val != t1.Val {
			return false
		}
		list.Push(t1.Left)
		list.Push(t2.Right)
		list.Push(t2.Left)
		list.Push(t1.Right)
	}
	return true
}

//复杂度分析：
//同上
