package _00_isSameTree

//100. 相同的树
//给定两个二叉树，编写一个函数来检验它们是否相同。
//
//如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
//
//示例 1:
//
//输入:       1         1
//           / \       / \
//          2   3     2   3
//
//[1,2,3],   [1,2,3]
//
//输出: true
//示例 2:
//
//输入:      1          1
//			/           \
//			2             2
//
//[1,2],     [1,null,2]
//
//输出: false
//示例 3:
//
//输入:       1         1
//			 / \       / \
//			2   1     1   2
//
//[1,2,1],   [1,1,2]
//
//输出: false
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归：
func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if q == nil || p == nil {
		return false
	}

	if q.Val != p.Val {
		return false
	}
	return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}

func check(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if q == nil || p == nil {
		return false
	}

	if q.Val != p.Val {
		return false
	}
	return true
}

//迭代：
func IsSameTree2(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, p)
	queue = append(queue, q)
	for len(queue) > 0 {
		pNode := queue[0]
		qNode := queue[1]
		if pNode.Val != qNode.Val {
			return false
		}
		if pNode.Left != nil && qNode.Left != nil {
			queue = append(queue, pNode.Left)
			queue = append(queue, qNode.Left)
		} else if pNode.Left == nil && qNode.Left != nil {
			return false
		} else if pNode.Left != nil && qNode.Left == nil {
			return false
		}
		if pNode.Right != nil && qNode.Right != nil {
			queue = append(queue, pNode.Right)
			queue = append(queue, qNode.Right)
		} else if pNode.Right == nil && qNode.Right != nil {
			return false
		} else if pNode.Right != nil && qNode.Right == nil {
			return false
		}
		queue = queue[2:]
	}
	return true
}
