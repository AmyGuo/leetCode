package _37_pathSum

//给定一个二叉树，它的每个结点都存放着一个整数值。
//
//找出路径和等于给定数值的路径总数。
//
//路径不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
//
//二叉树不超过1000个节点，且节点数值范围是 [-1000000,1000000] 的整数。
//
//示例：
//
//root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8
//
//10
///  \
//5   -3
/// \    \
//3   2   11
/// \   \
//3  -2   1
//
//返回 3。和等于 8 的路径有:
//
//1.  5 -> 3
//2.  5 -> 2 -> 1
//3.  -3 -> 11

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//方法一：递归法

//两个递归解决问题
//第一个递归：用于遍历每个结点
//第二个递归：从该节点开始向下找存在的路径个数

func PathSum(root *TreeNode, sum int) int {
	var total int = 0
	get(root, &total, sum)
	return total
}

//遍历所有节点
func get(root *TreeNode, total *int, sum int) {
	if root == nil {
		return
	}

	getNum(root, 0, sum, total)

	get(root.Left, total, sum)
	get(root.Right, total, sum)
}

//以当前节点为跟节点，向下计算存在的路径个数
func getNum(root *TreeNode, cur, sum int, total *int) {
	if root == nil {
		return
	}
	cur += root.Val
	if cur == sum {
		*total++
	}
	getNum(root.Left, cur, sum, total)
	getNum(root.Right, cur, sum, total)
}

//方法二：回溯法
//第一种做法很明显效率不够高，存在大量重复计算
//所以第二种做法，采取了类似于数组的前n项和的思路，比如sum[4] == sum[1]，那么1到4之间的和肯定为0
//
//对于树的话，采取DFS加回溯，每次访问到一个节点，把该节点加入到当前的pathSum中
//然后判断是否存在一个之前的前n项和，其值等于pathSum与sum之差
//如果有，就说明现在的前n项和，减去之前的前n项和，等于sum，那么也就是说，这两个点之间的路径和，就是sum
//
//最后要注意的是，记得回溯，把路径和弹出去
//dfs，其有两个重要的标志，也就是两个数组，一个用来标记该点是否被访问过，一个用来把该点放入数组，所以这两个标记是相辅相成的，一定同时出现；dfs就是随机选定一个起点将其标记为已经访问过的点，然后就是递归调用进行与其相邻的点的搜索，直到所有的点都被访问完。
func PathSum2(root *TreeNode, sum int) int {
	hash := make(map[int]int)
	cur := 0
	count := 0
	hash[0] = 1 //0 = pathSum-sum
	dfs(root, cur, sum, hash, &count)
	return count
}

func dfs(root *TreeNode, cur int, sum int, hash map[int]int, count *int) {
	if root == nil {
		return
	}
	cur += root.Val
	v, res := hash[cur-sum] //累加上到当前节点为止有多少条路径和等于pathSum的个数
	if res {
		*count = *count + v
	}
	v, res = hash[cur] //此处是计数到当前节点为止有多少条自上而下的节点路径和等于pathSum，并将其存入map
	if res {
		hash[cur] += 1
	} else {
		hash[cur] = 1
	}
	dfs(root.Left, cur, sum, hash, count)
	dfs(root.Right, cur, sum, hash, count)
	hash[cur] -= 1
}

//方法3：
var array = make([]int, 1000)

func PathSum3(root *TreeNode, sum int) int {
	return pathsum(root, sum, array, 0)
}

func pathsum(root *TreeNode, sum int, array []int, p int) int {
	if root == nil {
		return 0
	}
	tmp := root.Val
	var n int
	if root.Val == sum {
		n = 1
	} else {
		n = 0
	}
	for i := p - 1; i >= 0; i-- {
		tmp += array[i]
		if tmp == sum {
			n++
		}
	}
	array[p] = root.Val
	n1 := pathsum(root.Left, sum, array, p+1)
	n2 := pathsum(root.Right, sum, array, p+1)
	return n + n1 + n2
}
