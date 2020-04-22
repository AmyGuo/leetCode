package _0_climbStairs

//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
//注意：给定 n 是一个正整数。
//
//示例 1：
//
//输入： 2
//输出： 2
//解释： 有两种方法可以爬到楼顶。
//1.  1 阶 + 1 阶
//2.  2 阶
//示例 2：
//
//输入： 3
//输出： 3
//解释： 有三种方法可以爬到楼顶。
//1.  1 阶 + 1 阶 + 1 阶
//2.  1 阶 + 2 阶
//3.  2 阶 + 1 阶

//经典动态规划问题

//递归：f(i,n) = f(i+1,n) + f(i+2,n),   其中i是当前阶数，n是目标阶数
func ClimbStairs(n int) int {
	return cs(0, n)
}

func cs(i, n int) int {
	if i > n {
		return 0
	}

	if i == n {
		return 1
	}

	return cs(i+1, n) + cs(i+2, n)
}

//递归升级版
func ClimbStairs2(n int) int {
	arr := make([]int, n+1)
	return cs2(0, n, arr)
}

func cs2(i, n int, arr []int) int {
	if i > n {
		return 0
	}

	if i == n {
		return 1
	}

	if arr[i] > 0 {
		return arr[i]
	}

	arr[i] = cs2(i+1, n, arr) + cs2(i+2, n, arr)
	return arr[i]
}

//动态规划法：
//第 i 阶可以由以下两种方法得到：
//在第 (i-1)阶后向上爬一阶。
//在第 (i-2) 阶后向上爬 2 阶。
//所以到达第 ii 阶的方法总数就是到第 (i-1) 阶和第 (i-2) 阶的方法数之和。
//令 dp[i]表示能到达第 ii 阶的方法总数：
//dp[i]=dp[i-1]+dp[i-2]

func ClimbStairs3(n int) int {
	if n == 1 {
		return 1
	}
	arr := make([]int, n+1)
	arr[1] = 1
	arr[2] = 2
	for i := 3; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n]
}
