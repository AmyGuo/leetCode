package _5_threeSum

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
//
//注意：答案中不可以包含重复的三元组。
//
//
//
//示例：
//
//给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//满足要求的三元组集合为：
//[
//[-1, 0, 1],
//[-1, -1, 2]
//]

func ThreeSum(nums []int) [][]int {
	ret := make([][]int, 0)
	numMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			a := nums[i]
			b := nums[j]
			if _, ok := numMap[-a-b]; ok {
				ret = append(ret, []int{a, b, -a - b})
			}
		}
		numMap[i] = nums[i]
	}
	return ret
}
