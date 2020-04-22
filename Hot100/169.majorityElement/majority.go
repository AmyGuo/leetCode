package _69_majorityElement

//给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
//
//你可以假设数组是非空的，并且给定的数组总是存在多数元素。
//
//示例 1:
//
//输入: [3,2,3]
//输出: 3
//示例 2:
//
//输入: [2,2,1,1,1,2,2]
//输出: 2

func MajorityElement(nums []int) int {
	dict := make(map[int]int)
	for _, v := range nums {
		dict[v]++
		if dict[v] > len(nums)/2 {
			return v
		}
	}
	return 0
}

func MajorityElement2(nums []int) int {
	cur, count := 0, 0
	for _, v := range nums {
		if count == 0 {
			cur = v
		}
		if cur == v {
			count++
		} else {
			count--
		}
	}
	return cur
}
