package _5_threeSum

import "sort"

//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]


func threeSum(nums []int) [][]int {

	if len(nums) < 3 {
		return nil
	}

	sort.Ints(nums)
	numsLength := len(nums)

	var res [][]int
	for k := 0; k < numsLength - 2; k ++ {

		if nums[k] > 0 {
			break
		}

		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		left  := k + 1
		right := numsLength - 1

		for left < right {

			sum := nums[k] + nums[left] + nums[right]
			if sum > 0 {
				for right = right - 1; left < right && nums[right] == nums[right+1]; right -- {}
			}else if sum < 0 {
				for left = left + 1; left < right && nums[left] == nums[left-1]; left ++ {}
			}else {
				res = append(res, []int{nums[k], nums[left], nums[right]})
				for left = left + 1; left < right && nums[left] == nums[left-1]; left ++ {}
				for right = right - 1; left < right && nums[right] == nums[right+1]; right -- {}
			}
		}
	}

	return res
}


func threeSum2(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举 a
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first - 1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]
		// 枚举 b
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同
			if second > first + 1 && nums[second] == nums[second - 1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second] + nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[second] + nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}

