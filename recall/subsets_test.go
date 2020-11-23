package bitmap

import (
	"fmt"
	"testing"
)

//面试题 08.04. 幂集
//幂集。编写一种方法，返回某集合的所有子集。集合中不包含重复的元素。
//
//说明：解集不能包含重复的子集。
//
//示例:
//
//输入： nums = [1,2,3]
//输出：
//[
//[3],
//[1],
//[2],
//[1,2,3],
//[1,3],
//[2,3],
//[1,2],
//[]
//]

//回溯法

func subsets(nums []int) [][]int {
	res := [][]int{}
	var dfs func(path []int, index int)

	dfs = func(path []int, index int) {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)
		fmt.Println("res:", res)
		for i := index; i < len(nums); i++ {
			path = append(path, nums[i])
			fmt.Printf("path%+v,index:%d\n", path, i)
			dfs(path, i+1)
			path = path[:len(path)-1]
			fmt.Println("path return:", path)
		}
	}

	dfs([]int{}, 0)
	return res
}

func subsets1(nums []int) [][]int {
	// 结果集
	res := make([][]int, 0)
	// 临时栈
	stack := make([]int, 0)

	var search func(depth int)
	search = func(depth int) {
		if depth == len(nums) {
			t := make([]int, len(stack))
			copy(t, stack)
			res = append(res, t)
			return
		}
		// 递归
		stack = append(stack, nums[depth])
		search(depth + 1)
		// 回溯
		stack = stack[:len(stack)-1]
		search(depth + 1)
	}

	search(0)
	return res
}

func Test_subsets(t *testing.T) {
	fmt.Println(subsets([]int{1, 2, 3}))
}
