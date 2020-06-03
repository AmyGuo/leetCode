package _38_rangeSumBST

//938. 二叉搜索树的范围和
//给定二叉搜索树的根结点 root，返回 L 和 R（含）之间的所有结点的值的和。
//
//二叉搜索树保证具有唯一的值。
//
//
//
//示例 1：
//
//输入：root = [10,5,15,3,7,null,18], L = 7, R = 15
//输出：32
//示例 2：
//
//输入：root = [10,5,15,3,7,13,18,1,null,6], L = 6, R = 10
//输出：23
//
//
//提示：
//
//树中的结点数量最多为 10000 个。
//最终的答案保证小于 2^31。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归：
func RangeSumBST(root *TreeNode, L int, R int) int {
	var ret int
	dfs(root, &ret, L, R)
	return ret
}

func dfs(node *TreeNode, n *int, L, R int) {
	if node == nil {
		return
	}
	if node.Val < L {
		dfs(node.Right, n, L, R)
	} else if node.Val > R {
		dfs(node.Left, n, L, R)
	} else {
		*n += node.Val
		dfs(node.Left, n, L, R)
		dfs(node.Right, n, L, R)
	}
}

//迭代：
func RangeSumBST2(root *TreeNode, L int, R int) int {
	var ret int
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		temp := []*TreeNode{}
		for _, v := range nodes {
			if v.Val < L && v.Right != nil {
				temp = append(temp, v.Right)
			}
			if v.Val > R && v.Left != nil {
				temp = append(temp, v.Left)
			}

			if v.Val >= L && v.Val <= R {
				ret += v.Val
				if v.Left != nil {
					temp = append(temp, v.Left)
				}
				if v.Right != nil {
					temp = append(temp, v.Right)
				}
			}
		}
		nodes = temp
	}
	return ret
}
