package _01_findMod

//给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。
//
//假定 BST 有如下定义：
//
//结点左子树中所含结点的值小于等于当前结点的值
//结点右子树中所含结点的值大于等于当前结点的值
//左子树和右子树都是二叉搜索树
//例如：
//给定 BST [1,null,2,2],
//
//   1
//    \
//     2
//    /
//   2
//返回[2].
//
//提示：如果众数超过1个，不需考虑输出顺序
//
//进阶：你可以不使用额外的空间吗？（假设由递归产生的隐式调用栈的开销不被计算在内）

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var max int     //最大值
var res []int   //结果
var cur int     //当前
var counter int //当前计数

func FindMode(root *TreeNode) []int {
	res, max, cur, counter = []int{}, 1, 0, 0
	dfs(root)
	return res
}

func dfs(root *TreeNode) {
	if root != nil {
		dfs(root.Left)
		if root.Val != cur {
			counter = 0
		}
		counter++
		if max < counter {
			max = counter
			res = []int{root.Val}
		} else if max == counter {
			res = append(res, root.Val)
		}
		cur = root.Val
		dfs(root.Right)
	}
}
