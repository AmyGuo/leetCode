package _59_MaxDepth

//N叉树的最大深度
//给定一个 N 叉树，找到其最大深度。
//
//最大深度是指从根节点到最远叶子节点的最长路径上的节点总数。
//
//例如，给定一个 3叉树 :
//					1
//				 /  |  \
//				3    2   4
//			  /	 \
//			 5	  6
//
//我们应返回其最大深度，3。
//
//说明:
//
//树的深度不会超过 1000。
//树的节点总不会超过 5000。

type Node struct {
	Val      int
	Children []*Node
}

func MaxDepth(root *Node) int {
	if root == nil {
		return 0
	} else if root.Children == nil {
		return 1
	} else {
		max := 0
		for _, v := range root.Children {
			tmp := MaxDepth(v) + 1
			if max < tmp {
				max = tmp
			}
		}
		return max
	}
}

func MaxDepth2(root *Node) int {
	if root == nil {
		return 0
	}
	list := []*Node{root}
	depth := 0
	for len(list) > 0 {
		length := len(list)
		for length > 0 {
			length--
			for _, v := range list[0].Children {
				list = append(list, v)
			}
			list = list[1:]
		}
		depth++
	}
	return depth
}
