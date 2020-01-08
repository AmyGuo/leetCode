package _04_maxDepth

//给定一个二叉树，找出其最大深度。
//
//二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
//
//说明: 叶子节点是指没有子节点的节点。
//
//示例：
//给定二叉树 [3,9,20,null,null,15,7]，
//
//3
/// \
//9  20
///  \
//15   7
//返回它的最大深度 3 。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := MaxDepth(root.Left)
	right := MaxDepth(root.Right)
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

//迭代
type Node struct {
	T     *TreeNode
	Depth int
}

func MaxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 1
	stack := []*Node{&Node{T: root, Depth: 1}}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Depth > ans {
			ans = node.Depth
		}
		if node.T.Right != nil {
			stack = append(stack, &Node{T: node.T.Right, Depth: node.Depth + 1})
		}
		if node.T.Left != nil {
			stack = append(stack, &Node{T: node.T.Left, Depth: node.Depth + 1})
		}
	}
	return ans
}
