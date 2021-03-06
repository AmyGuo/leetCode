package dynamicPlanning

import "fmt"

/*
面试题 17.16. 按摩师
一个有名的按摩师会收到源源不断的预约请求，每个预约都可以选择接或不接。在每次预约服务之间要有休息时间，因此她不能接受相邻的预约。
给定一个预约请求序列，替按摩师找到最优的预约集合（总预约时间最长），返回总的分钟数。

注意：本题相对原题稍作改动



示例 1：
输入： [1,2,3,1]
输出： 4
解释： 选择 1 号预约和 3 号预约，总时长 = 1 + 3 = 4。

示例 2：
输入： [2,7,9,3,1]
输出： 12
解释： 选择 1 号预约、 3 号预约和 5 号预约，总时长 = 2 + 9 + 1 = 12。

示例 3：
输入： [2,1,4,5,3,1,1,3]
输出： 12
解释： 选择 1 号预约、 3 号预约、 5 号预约和 8 号预约，总时长 = 2 + 4 + 3 + 3 = 12。
*/

//此题和之前的“抢劫”题很像，是一个动态规划的典型题目。
//对于每个预约，只能选择接或不接。
//如果接，则上一个状态必须是不接。
//如果不接，上一个状态可以是接，也可以不接。
//
//每个点有两个状态（接，不接），因此需要定义一个二维数组用于存储。
//
//因此，dp方程是：
//dp[k][接] = dp[k-1][不接] + nums[k]
//dp[k][不接] = max(dp[k-1][接],dp[k-1][不接])

func massage(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2) // [0]定义为接，[1]定义为不接
	}
	dp[0][0] = nums[0]
	dp[0][1] = 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][1] + nums[i]
		dp[i][1] = max(dp[i-1][0], dp[i-1][1])
	}
	return max(dp[n-1][0], dp[n-1][1])
}

func ExampleMassage() {
	fmt.Println(massage([]int{1, 2, 3, 1}))
	fmt.Println(massage([]int{2, 7, 9, 3, 1}))
	fmt.Println(massage([]int{2, 1, 4, 5, 3, 1, 1, 3}))
	//Output:
	//4
	//12
	//12
}
