package _57_treePaths

import "strconv"

//给定一个二叉树，返回所有从根节点到叶子节点的路径。
//
//说明: 叶子节点是指没有子节点的节点。
//
//示例:
//
//输入:
//
//  1
///   \
//2     3
//  \
//   5
//
//输出: ["1->2->5", "1->3"]
//
//解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BinaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	ret := make([]string, 0)
	s := ""
	getNext(root, &ret, s)
	return ret
}

func getNext(node *TreeNode, ret *[]string, s string) string {
	if node.Left == nil && node.Right == nil {
		s += strconv.Itoa(node.Val)
		*ret = append(*ret, s)
		return ""
	}

	s += strconv.Itoa(node.Val)
	s += "->"
	if node.Left != nil {
		getNext(node.Left, ret, s)
	}
	if node.Right != nil {
		getNext(node.Right, ret, s)
	}
	return s
}

//  1
///   \
//2     3
//  \
//   5
func BinaryTreePaths2(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	str := strconv.Itoa(root.Val)
	res := []string{}

	if root.Left != nil {
		strs := BinaryTreePaths2(root.Left)
		for _, v := range strs {
			res = append(res, str+"->"+v)
		}
	}
	if root.Right != nil {
		strs := BinaryTreePaths2(root.Right)
		for _, v := range strs {
			res = append(res, str+"->"+v)
		}
	}
	if len(res) == 0 {
		return []string{str}
	}
	return res
}
