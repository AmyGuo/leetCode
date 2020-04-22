package _60_subarraySum

//给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。
//
//示例 1 :
//
//输入:nums = [1,1,1], k = 2
//输出: 2 , [1,1] 与 [1,1] 为两种不同的情况。
//说明 :
//
//数组的长度为 [1, 20,000]。
//数组中元素的范围是 [-1000, 1000] ，且整数 k 的范围是 [-1e7, 1e7]

//暴力法
func SubarraySum(nums []int, k int) int {
	count := 0
	for start := 0; start < len(nums); start++ {
		for end := start + 1; end <= len(nums); end++ {
			sum := 0
			for i := start; i < end; i++ {
				sum += nums[i]
			}
			if sum == k {
				count++
			}
		}
	}
	return count
}

//通过累积和来计算
func SubarraySum2(nums []int, k int) int {
	count := 0
	sum := make([]int, len(nums)+1)
	sum[0] = 0
	for i := 1; i <= len(nums); i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}
	for start := 0; start < len(nums); start++ {
		for end := start + 1; end <= len(nums); end++ {
			if sum[end]-sum[start] == k {
				count++
			}
		}
	}
	return count
}

//优化累积和方式，不适用额外的内存空间
func SubarraySum3(nums []int, k int) int {
	count := 0
	for start := 0; start < len(nums); start++ {
		sum := 0
		for end := start; end < len(nums); end++ {
			sum += nums[end]
			if sum == k {
				count++
			}
		}
	}
	return count
}

func SubarraySum4(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 && nums[0] != k {
		return 0
	}

	count, sum := 0, 0
	m := map[int]int{0: 1}
	for _, v := range nums {
		sum += v
		if temp, ok := m[sum-k]; ok {
			count += temp
		}

		m[sum]++
	}
	return count
}
