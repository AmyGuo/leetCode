package _89_preOrder

//590. N叉树的前序遍历
//给定一个 N 叉树，返回其节点值的后序遍历。
//
//例如，给定一个 3叉树 :
//					1
//				 /  |  \
//				3    2   4
//			  /	 \
//			 5	  6
//返回其前序遍历: [1.3.5.6.2.4].
//说明: 递归法很简单，你可以使用迭代法完成此题吗?

type Node struct {
	Val      int
	Children []*Node
}

func PreOrder(root *Node) []int {
	if root == nil {
		return nil
	}
	ret := make([]int, 0)
	dfs(root, &ret)
	return ret
}

func dfs(node *Node, ret *[]int) {
	*ret = append(*ret, node.Val)

	for _, v := range node.Children {
		dfs(v, ret)
	}
}

func PreOrder2(root *Node) []int {
	if root == nil {
		return nil
	}
	stack := []*Node{root}
	ret := make([]int, 0)
	for len(stack) > 0 {
		root, stack = stack[len(stack)-1], stack[:len(stack)-1]
		ret = append(ret, root.Val)
		for i := len(root.Children) - 1; i >= 0; i-- {
			if child := root.Children[i]; child != nil {
				stack = append(stack, child)
			}
		}
	}
	return ret
}
