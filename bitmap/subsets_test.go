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

func subsets(nums []int) [][]int {
	return nil
}

func Test_subsets(t *testing.T) {
	fmt.Println(subsets([]int{1, 2, 3}))
}
