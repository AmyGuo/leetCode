package recall

import (
	"fmt"
	"testing"
)

/*
输入一个不包含重复数字的数组 nums，返回这些数字的全部排列。
比如说输入数组 [1,2,3]，输出结果应该如下，顺序无所谓，不能有重复：
[ [1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], [3,2,1] ]
*/

func premute(nums []int) [][]int {
	res := make([][]int, 0)
	var dfs func(used map[int]bool, path []int)
	dfs = func(used map[int]bool, path []int) {
		if len(path) == len(nums) {
			res = append(res, append([]int(nil), path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[nums[i]] {
				continue
			}
			path = append(path, nums[i])
			used[nums[i]] = true
			dfs(used, path)
			path = path[:len(path)-1]
			used[nums[i]] = false
		}
	}

	m := make(map[int]bool)
	dfs(m, []int{})
	return res
}

func TestPremute(t *testing.T) {
	fmt.Println(premute([]int{1, 2, 3}))
	//[[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]
}
