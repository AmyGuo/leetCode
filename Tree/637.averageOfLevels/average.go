package _37_averageOfLevels

//给定一个非空二叉树, 返回一个由每层节点平均值组成的数组.
//
//示例 1:
//
//输入:
//3
/// \
//9  20
///  \
//15   7
//输出: [3, 14.5, 11]
//解释:
//第0层的平均值是 3,  第1层是 14.5, 第2层是 11. 因此返回 [3, 14.5, 11].

//一般树的操作都可分为广度和深度的遍历，即迭代和递归两种方法
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//分析：使用两个数组，一个记录每一层的数字之和，一个记录每一层的数字个数

//递归法：
func AverageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	sum := make([]float64, 0)
	num := make([]float64, 0)
	ret := make([]float64, 0)
	average(root, 0, &sum, &num)
	for i, s := range sum {
		ret = append(ret, s/num[i])
	}
	return ret
}

func average(node *TreeNode, i int, sum, num *[]float64) {
	if node == nil {
		return
	}

	if i < len(*sum) {
		(*sum)[i] += float64(node.Val)
		(*num)[i]++

	} else {
		*sum = append(*sum, float64(node.Val))
		*num = append(*num, 1)
	}

	average(node.Left, i+1, sum, num)
	average(node.Right, i+1, sum, num)
}

//迭代法：(层次遍历)
func AverageOfLevels2(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}

	nodes := []*TreeNode{root}
	res := make([]float64, 0)
	for len(nodes) > 0 {
		sum := 0
		curNodes := []*TreeNode{}
		for _, v := range nodes {
			sum += v.Val
			if v.Left != nil {
				curNodes = append(curNodes, v.Left)
			}
			if v.Right != nil {
				curNodes = append(curNodes, v.Right)
			}
		}
		res = append(res, float64(sum)/float64(len(nodes)))
		nodes = curNodes
	}
	return res
}
