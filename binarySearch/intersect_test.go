package binarySearch

import "fmt"

/*
350. 两个数组的交集 II
给定两个数组，编写一个函数来计算它们的交集。



示例 1：
输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2,2]


示例 2:
输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[4,9]


说明：

输出结果中每个元素出现的次数，应与元素在两个数组中出现次数的最小值一致。
我们可以不考虑输出结果的顺序。
进阶：

如果给定的数组已经排好序呢？你将如何优化你的算法？
如果 nums1 的大小比 nums2 小很多，哪种方法更优？
如果 nums2 的元素存储在磁盘上，内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
*/

//func intersect(nums1 []int, nums2 []int) []int {
//	if len(nums1) > len(nums2) {
//		return intersect(nums2, nums1)
//	}
//	m := map[int]int{}
//	for _, num := range nums1 {
//		m[num]++
//	}
//
//	intersection := []int{}
//	for _, num := range nums2 {
//		if m[num] > 0 {
//			intersection = append(intersection, num)
//			m[num]--
//		}
//	}
//	return intersection
//}

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}

	m := make(map[int]int)
	for _, v := range nums1 {
		m[v]++
	}

	ret := make([]int, 0)
	for _, num := range nums2 {
		if m[num] > 0 {
			ret = append(ret, num)
			m[num]--
		}
	}
	return ret
}

func ExampleIntersect() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
	//Output:
	//[2 2]
	//[9 4]
}
