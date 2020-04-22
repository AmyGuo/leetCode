package _48_findDisappearedNumbers

//给定一个范围在  1 ≤ a[i] ≤ n ( n = 数组大小 ) 的 整型数组，数组中的元素一些出现了两次，另一些只出现一次。
//
//找到所有在 [1, n] 范围之间没有出现在数组中的数字。
//
//您能在不使用额外空间且时间复杂度为O(n)的情况下完成这个任务吗? 你可以假定返回的数组不算在额外空间内。
//
//示例:
//
//输入:
//[4,3,2,7,8,2,3,1]
//
//输出:
//[5,6]

//hash法：
func FindDisappearedNumbers(nums []int) []int {
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}
	ret := make([]int, 0)
	i := 1
	for i <= len(nums) {
		if _, ok := m[i]; !ok {
			ret = append(ret, i)
		}
		i++
	}
	return ret
}

//鸽巢理论：
//遍历输入数组的每个元素一次。
//我们将把 |nums[i]|-1 索引位置的元素标记为负数。即 nums[|nums[i] |- 1] ×−1 。
//然后遍历数组，若当前数组元素 nums[i] 为负数，说明我们在数组中存在数字 i+1。

func FindDisappearedNumbers2(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		newIndex := 0
		if nums[i] > 0 {
			newIndex = nums[i] - 1
		} else {
			newIndex = -nums[i] - 1
		}
		if nums[newIndex] > 0 {
			nums[newIndex] *= -1
		}
	}

	ret := make([]int, 0)
	for i := 1; i <= len(nums); i++ {
		if nums[i-1] > 0 {
			ret = append(ret, i)
		}
	}
	return ret
}
