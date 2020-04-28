package _93_isCousins

//在二叉树中，根节点位于深度 0 处，每个深度为 k 的节点的子节点位于深度 k+1 处。
//
//如果二叉树的两个节点深度相同，但父节点不同，则它们是一对堂兄弟节点。
//
//我们给出了具有唯一值的二叉树的根节点 root，以及树中两个不同节点的值 x 和 y。
//
//只有与值 x 和 y 对应的节点是堂兄弟节点时，才返回 true。否则，返回 false。
//
//
//
//示例 1：
//    1
//   /  \
//  2    3
// /
//4
//输入：root = [1,2,3,4], x = 4, y = 3
//输出：false

//示例 2：
//
////        1
//         / \
//        2    3
//         \    \
//          4    5
//输入：root = [1,2,3,null,4,null,5], x = 5, y = 4
//输出：true
//示例 3：
//
////        1
//         / \
//        2    3
//         \
//          4
//输入：root = [1,2,3,null,4], x = 2, y = 3
//输出：false
//
//
//提示：
//
//二叉树的节点数介于 2 到 100 之间。
//每个节点的值都是唯一的、范围为 1 到 100 的整数。
//

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//1. 判断两个节点的深度
//2. 两个节点的父节点是否一致
//迭代法：
func IsCousins(root *TreeNode, x int, y int) bool {
	depthMap := make(map[int]int)
	fatherMap := make(map[int]*TreeNode)
	dfs(root, nil, 0, &depthMap, &fatherMap)
	return depthMap[x] == depthMap[y] && fatherMap[x] != fatherMap[y]
}

func dfs(node, parent *TreeNode, layer int, depthMap *map[int]int, fatherMap *map[int]*TreeNode) {
	if node == nil {
		return
	}
	(*depthMap)[node.Val] = layer
	(*fatherMap)[node.Val] = parent

	dfs(node.Left, node, layer+1, depthMap, fatherMap)
	dfs(node.Right, node, layer+1, depthMap, fatherMap)
}

//递归法：
func IsCousins2(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}

	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		var vx, vy bool
		var px, py int
		temp := []*TreeNode{}
		for _, v := range nodes {
			if v.Left != nil {
				temp = append(temp, v.Left)
				if v.Left.Val == x {
					vx = true
					px = v.Val
				}
				if v.Left.Val == y {
					vy = true
					py = v.Val
				}
			}

			if v.Right != nil {
				temp = append(temp, v.Right)
				if v.Right.Val == x {
					vx = true
					px = v.Val
				}
				if v.Right.Val == y {
					vy = true
					py = v.Val
				}
			}

		}
		if vx && vy {
			return px != py
		}
		nodes = temp
	}
	return false
}
