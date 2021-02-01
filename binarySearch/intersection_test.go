package binarySearch

import "fmt"

/*
349. 两个数组的交集
给定两个数组，编写一个函数来计算它们的交集。


示例 1：
输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2]


示例 2：
输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[9,4]


说明：

输出结果中的每个元素一定是唯一的。
我们可以不考虑输出结果的顺序。
*/

func intersection(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]bool)
	for _, v := range nums1 {
		if _, ok := m1[v]; ok {
			continue
		}
		m1[v] = true
	}

	ret := make([]int, 0)
	retM := make(map[int]bool)
	for _, v := range nums2 {
		if m1[v] && !retM[v] {
			retM[v] = true
			ret = append(ret, v)
		}
	}
	return ret
}

func ExampleInterSection() {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
	//Output:
	//[2]
	//[9 4]
}
