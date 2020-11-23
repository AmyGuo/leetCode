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

//数组中的每一个数字都有选和不选两种状态，我们可以用0和1表示，0表示不选，1表示选择。如果数组的长度是n，那么子集的数量就是2^n。比如数组长度是3，就有8种可能，分别是
//
//[0，0，0]
//
//[0，0，1]
//
//[0，1，0]
//
//[0，1，1]
//
//[1，0，0]
//
//[1，0，1]
//
//[1，1，0]
//
//[1，1，1]

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	length := uint(1 << uint(len(nums)))
	for i := uint(0); i < length; i++ { //遍历从0到length中间的所有数字，根据数字中1的位置来找子集
		list := make([]int, 0)
		for j := uint(0); j < length; j++ { //遍历该数字的每一位
			if ((i >> j) & 1) == 1 { //如果数字i的某一个位置是1，就把数组中对应的数字添加到集合
				list = append(list, nums[j])
			}
		}
		res = append(res, list)
	}
	return res
}

func subsets1(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}
	result := subsets(nums[1:])
	for _, v := range result {
		result = append(result, append([]int{nums[0]}, v...))
	}
	return result
}

func Test_subsets(t *testing.T) {
	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsets1([]int{1, 2, 3}))
}
