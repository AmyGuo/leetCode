package recall

import (
	"fmt"
	"testing"
)

/*
39. 组合总和
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。
示例 1：

输入：candidates = [2,3,6,7], target = 7,
所求解集为：
[
  [7],
  [2,2,3]
]
示例 2：

输入：candidates = [2,3,5], target = 8,
所求解集为：
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]


提示：

1 <= candidates.length <= 30
1 <= candidates[i] <= 200
candidate 中的每个元素都是独一无二的。
1 <= target <= 500
*/

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}

	res := make([][]int, 0)
	temp := target

	var backTrack func(index, temp int, pth []int)
	backTrack = func(index, temp int, pth []int) {
		if index == len(candidates) {
			return
		}

		if temp == 0 {
			res = append(res, pth)
			pth = []int{}
			return
		}

		backTrack(index+1, temp, pth)

		if temp-candidates[index] >= 0 {
			pth = append(pth, candidates[index])
			backTrack(index, temp-candidates[index], pth)
			pth = pth[:len(pth)-1]
		}
	}

	backTrack(0, temp, []int{})

	return res
}

func Test_combinationSum(t *testing.T) {
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}
