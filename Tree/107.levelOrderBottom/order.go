package _07_levelOrderBottom

//给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
//
//例如：
//给定二叉树 [3,9,20,null,null,15,7],
//
//3
/// \
//9  20
///  \
//15   7
//返回其自底向上的层次遍历为：
//
//[
//[15,7],
//[9,20],
//[3]
//]

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	temp := make([][]int, 0)
	getNextLayer(root, 0, &temp)

	//res := make([][]int, 0)
	//for i := len(temp) - 1; i >= 0; i-- {
	//	res = append(res, temp[i])
	//}
	//return res

	l := len(temp)
	end := l / 2
	for i := 0; i < end; i++ {
		temp[i], temp[l-i-1] = temp[l-i-1], temp[i]
	}
	return temp
}

func getNextLayer(node *TreeNode, layer int, res *[][]int) {
	if node == nil {
		return
	}
	if layer < len(*res) {
		(*res)[layer] = append((*res)[layer], node.Val)
	} else {
		(*res) = append((*res), []int{node.Val})
	}
	getNextLayer(node.Left, layer+1, res)
	getNextLayer(node.Right, layer+1, res)
}

func LevelOrderBottom2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		var temp []*TreeNode
		var vals []int
		for _, v := range nodes {
			vals = append(vals, v.Val)
			if v.Left != nil {
				temp = append(temp, v.Left)
			}
			if v.Right != nil {
				temp = append(temp, v.Right)
			}
		}
		res = append(res, vals)
		nodes = temp
	}
	result := make([][]int, 0)
	for i := len(res) - 1; i >= 0; i-- {
		result = append(result, res[i])
	}
	return result

}
