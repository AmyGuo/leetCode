package _94_findTargetSumWays

//给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。对于数组中的任意一个整数，你都可以从 + 或 -中选择一个符号添加在前面。
//
//返回可以使最终数组和为目标数 S 的所有添加符号的方法数。
//
//示例 1:
//
//输入: nums: [1, 1, 1, 1, 1], S: 3
//输出: 5
//解释:
//
//-1+1+1+1+1 = 3
//+1-1+1+1+1 = 3
//+1+1-1+1+1 = 3
//+1+1+1-1+1 = 3
//+1+1+1+1-1 = 3
//
//一共有5种方法让最终目标和为3。
//注意:
//
//数组非空，且长度不会超过20。
//初始的数组的和不会超过1000。
//保证返回的最终结果能被32位整数存下。

//枚举法：使用递归列出所有的情况
func FindTargetSumWays(nums []int, S int) int {
	count := 0
	calculate(nums, 0, 0, S, &count)
	return count
}

func calculate(nums []int, i, sum, S int, count *int) {
	if i == len(nums) {
		if sum == S {
			*count++
		}
	} else {
		calculate(nums, i+1, sum+nums[i], S, count)
		calculate(nums, i+1, sum-nums[i], S, count)
	}
}

//动态规划法  背包问题？？？
//我们用 dp[i][j] 表示用数组中的前 i 个元素，组成和为 j 的方案数。考虑第 i 个数 nums[i]，它可以被添加 + 或 -，因此状态转移方程如下：
//dp[i][j] = dp[i - 1][j - nums[i]] + dp[i - 1][j + nums[i]]
func FindTargetSumWays2(nums []int, S int) int {

}
