package base

import (
	"fmt"
)

//二叉树的遍历

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func (node *TreeNode) Print() {
	fmt.Print(node.Value, " ")
}

func (node *TreeNode) SetValue(v int) {
	if node == nil {
		fmt.Println("setting value to nil.node ignored.")
		return
	}
	node.Value = v
}

//前序遍历
func (node *TreeNode) PreOrder() {
	if node == nil {
		return
	}
	node.Print()
	node.Left.PreOrder()
	node.Right.PreOrder()
}

//中序遍历
func (node *TreeNode) MiddleOrder() {
	if node == nil {
		return
	}
	node.Left.MiddleOrder()
	node.Print()
	node.Right.MiddleOrder()
}

//后序遍历
func (node *TreeNode) PostOrder() {
	if node == nil {
		return
	}
	node.Left.PostOrder()
	node.Right.PostOrder()
	node.Print()
}

//层次遍历(广度优先遍历)
func (node *TreeNode) BreadthFirstSearch() {
	if node == nil {
		return
	}
	result := []int{}
	nodes := []*TreeNode{node}
	for len(nodes) > 0 {
		curTreeNode := nodes[0]
		nodes = nodes[1:]
		result = append(result, curTreeNode.Value)
		if curTreeNode.Left != nil {
			nodes = append(nodes, curTreeNode.Left)
		}
		if curTreeNode.Right != nil {
			nodes = append(nodes, curTreeNode.Right)
		}
	}
	for _, v := range result {
		fmt.Print(v, " ")
	}
}

//层数(递归实现)
//对任意一个子树的根节点来说，它的深度=左右子树深度的最大值+1
func (node *TreeNode) Layers() int {
	if node == nil {
		return 0
	}
	leftLayers := node.Left.Layers()
	rightLayers := node.Right.Layers()
	if leftLayers > rightLayers {
		return leftLayers + 1
	} else {
		return rightLayers + 1
	}
}

//层数(非递归实现)
//借助队列，在进行按层遍历时，记录遍历的层数即可
func (node *TreeNode) LayersByQueue() int {
	if node == nil {
		return 0
	}
	layers := 0
	nodes := []*TreeNode{node}
	for len(nodes) > 0 {
		layers++
		size := len(nodes) //每层的节点数
		count := 0
		for count < size {
			count++
			curTreeNode := nodes[0]
			nodes = nodes[1:]
			if curTreeNode.Left != nil {
				nodes = append(nodes, curTreeNode.Left)
			}
			if curTreeNode.Right != nil {
				nodes = append(nodes, curTreeNode.Right)
			}
		}
	}
	return layers
}

func CreateTreeNode(v int) *TreeNode {
	return &TreeNode{Value: v}
}

//根据输入的数据构建二叉树
func BuildTreeNode(nums []interface{}) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	ret := new(TreeNode)
	build(ret, nums, 0)
	return ret
}

func build(tree *TreeNode, nums []interface{}, index int) {
	if index > len(nums)-1 {
		return
	}

	if val, ok := nums[index].(int); ok {
		tree.Value = val
		tree.Left = new(TreeNode)
		tree.Right = new(TreeNode)
		build(tree.Left, nums, 2*index+1)
		build(tree.Right, nums, 2*index+2)
	} else {
		tree.Left = nil
		tree.Right = nil
	}
}
