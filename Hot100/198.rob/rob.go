package _98_rob

//打家劫舍
//你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
//
//给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。
//
//示例 1:
//
//输入: [1,2,3,1]
//输出: 4
//解释: 偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
//     偷窃到的最高金额 = 1 + 3 = 4 。
//示例 2:
//
//输入: [2,7,9,3,1]
//输出: 12
//解释: 偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
//     偷窃到的最高金额 = 2 + 9 + 1 = 12 。
//

//动态规划类型的题目，重点是找到状态转移方程

//采坑点：不是相邻的数相加就是最大值，间隔数可以为1，也可以>1
func Rob(nums []int) int {
	evenSum := 0
	oddSum := 0
	for i, v := range nums {
		if i%2 == 0 {
			evenSum += v
			if evenSum < oddSum {
				evenSum = oddSum
			}
		} else if i%2 != 0 {
			oddSum += v
			if evenSum > oddSum {
				oddSum = evenSum
			}
		}
	}
	if evenSum > oddSum {
		return evenSum
	} else {
		return oddSum
	}
}

//动态规划法：
//转移方程式：
//dp[n+1] = max(dp[n],dp[n-1]+num)

func Rob2(nums []int) int {
	cur := 0
	pre := 0
	for _, v := range nums {
		if pre+v > cur {
			cur, pre = pre+v, cur
		} else {
			cur, pre = cur, cur
		}
		//cur, pre = math.Max(pre+v, cur), cur
	}
	return cur
}

func Rob3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//状态定义：
//
//设动态规划列表 dp ，dp[i] 代表前 i 个房子在满足条件下的能偷窃到的最高金额。
//转移方程：
//
//设： 有 n 个房子，前 n 间能偷窃到的最高金额是 dp[n] ，前 n-1 间能偷窃到的最高金额是dp[n−1] ，此时向这些房子后加一间房，此房间价值为 num ；
//
//加一间房间后： 由于不能抢相邻的房子，意味着抢第 n+1 间就不能抢第 n 间；那么前n+1 间房能偷取到的最高金额]dp[n+1] 一定是以下两种情况的 较大值 ：
//
//不抢第 n+1 个房间，因此等于前 n 个房子的最高金额，即 dp[n+1]=dp[n] ；
//抢第 n+1 个房间，此时不能抢第 n 个房间；因此等于前 n-1 个房子的最高金额加上当前房间价值，即 dp[n+1]=dp[n−1]+num ；
//
//最终的转移方程： dp[n+1] = max(dp[n],dp[n-1]+num)
//
//初始状态：
//
//前 0 间房子的最大偷窃价值为 0 ，即 dp[0] = 0。
//返回值：
//
//返回 dp 列表最后一个元素值，即所有房间的最大偷窃价值。
//简化空间复杂度：
//
//我们发现 dp[n] 只与 dp[n-1] 和 dp[n-2] 有关系，因此我们可以设两个变量 cur和 pre 交替记录，将空间复杂度降到 O(1) 。
//复杂度分析：
//时间复杂度 O(N) ： 遍历 nums 需要线性时间；
//空间复杂度 O(1) ： cur和 pre 使用常数大小的额外空间。
//
