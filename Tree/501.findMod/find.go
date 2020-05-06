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

//查找一组（无序的）数中的众数，无非就是遍历数组，统计每个数字的出现频率。那么二叉搜索树可以提供什么信息呢？
//二叉搜索树的左右子结点和父结点之间有大小关系的限制，且二叉树的中序遍历是升序的。由此，问题可以转变成在一组升序排列的数中查找众数。
//
//既然数字是升序的，就可以遍历一次完成统计。
//借助三个变量，一个记录当前数字，一个记录当前数字的频率，一个记录上一个添加到List的数字的频率。
//
//每访问一个数字，就将这个数字的出现次数加一，如果当前出现次数等于上一个添加到List的数字的出现次数，则在结果List中添加这个数字；
//如果当前出现次数大于上一个添加到List的数字的出现次数，则清空结果List再添加这个数字，并把上一个添加到List数字的出现次数更新为当前出现次数。
//
//还需注意两点：
//若是首次访问，需要将上一个添加到List的数字的频率设为1；
//每遇到一个新数字，则将当前数字的出现次数清零。

func FindMode(root *TreeNode) []int {
	var max int     //最大值
	var cur int     //当前
	var counter int //当前计数
	var res []int   //结果
	dfs(root, &max, &cur, &counter, &res)
	return res
}

func dfs(node *TreeNode, max, cur, counter *int, res *[]int) {
	if node == nil {
		return
	}

	dfs(node.Left, max, cur, counter, res)
	if node.Val != *cur {
		*counter = 0
	}
	*counter++
	if *max < *counter {
		*max = *counter
		*res = []int{node.Val}
	} else if *max == *counter {
		*res = append(*res, node.Val)
	}
	*cur = node.Val
	dfs(node.Right, max, cur, counter, res)
}
