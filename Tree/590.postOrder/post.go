package _90_postOrder

//590. N叉树的后序遍历
//给定一个 N 叉树，返回其节点值的后序遍历。
//
//例如，给定一个 3叉树 :
//					1
//				 /  |  \
//				3    2   4
//			  /	 \
//			 5	  6
//返回其后序遍历: [5,6,3,2,4,1].
//说明: 递归法很简单，你可以使用迭代法完成此题吗?

type Node struct {
	Val      int
	Children []*Node
}

//递归：
func PostOrder(root *Node) []int {
	if root == nil {
		return nil
	}
	ret := make([]int, 0)
	dfs(root, &ret)
	return ret
}

func dfs(node *Node, ret *[]int) {
	for _, v := range node.Children {
		dfs(v, ret)
	}
	*ret = append(*ret, node.Val)
}

//迭代：
func PostOrder2(root *Node) []int {
	if root == nil {
		return nil
	}
	ret := make([]int, 0)
	nodes := []*Node{root}
	for len(nodes) > 0 {
		last := nodes[len(nodes)-1]
		temp := nodes[:len(nodes)-1]
		ret = append([]int{last.Val}, ret...)
		for _, v := range last.Children {
			temp = append(temp, v)
		}

		nodes = temp
	}
	return ret
}
