package greedy

import "fmt"

/*
1005. K 次取反后最大化的数组和
给定一个整数数组 A，我们只能用以下方法修改该数组：我们选择某个索引 i 并将 A[i] 替换为 -A[i]，然后总共重复这个过程 K 次。（我们可以多次选择同一个索引 i。）

以这种方式修改数组后，返回数组可能的最大和。



示例 1：

输入：A = [4,2,3], K = 1
输出：5
解释：选择索引 (1,) ，然后 A 变为 [4,-2,3]。
示例 2：

输入：A = [3,-1,0,2], K = 3
输出：6
解释：选择索引 (1, 2, 2) ，然后 A 变为 [3,1,0,2]。
示例 3：

输入：A = [2,-3,-1,5,-4], K = 2
输出：13
解释：选择索引 (1, 4) ，然后 A 变为 [2,3,-1,5,4]。


提示：
1 <= A.length <= 10000
1 <= K <= 10000
-100 <= A[i] <= 100
*/

func ExampleLargestSumAfterKNegations() {
	fmt.Println(largestSumAfterKNegations([]int{4, 2, 3}, 1))
	fmt.Println(largestSumAfterKNegations([]int{3, -1, 0, 2}, 3))
	fmt.Println(largestSumAfterKNegations([]int{2, -3, -1, 5, -4}, 2))
	fmt.Println(largestSumAfterKNegations([]int{1, 2, 3, 4, 5}, 1))
	//Output:
	//5
	//6
	//13
	//13
}

func largestSumAfterKNegations(A []int, K int) int {

	var s = [201]int{}
	sum := 0
	for _, value := range A {
		s[value+100]++
	}
	temp := 0
	for key, value := range s {
		if value > 0 {
			if K > 0 && key < 100 {
				if K >= value {
					sum = sum - (key-100)*value
					K -= value
				} else {
					sum = sum - (key-100)*K + (key-100)*(value-K)
					K = 0
				}
			} else if K > 0 && key >= 100 { //todo:还没有看懂
				fmt.Println(key, K, K&1)
				if K&1 == 0 { //奇数
					sum += (key - 100) * value
				} else {
					if -(temp - 100) < key-100 {
						sum = sum + (temp-100)*2 + (key-100)*value
					} else {
						sum = sum - (key-100)*1 + (key-100)*(value-1)
					}
				}
				K = 0
			} else {
				sum += (key - 100) * value
			}
			temp = key
		}
	}
	return sum
}
