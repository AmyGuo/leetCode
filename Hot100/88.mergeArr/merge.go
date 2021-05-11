package Hot100

import "fmt"

//合并两个有序数组

//func merge(nums1 []int, m int, nums2 []int, n int) []int {
//for p1, p2, tail := m-1, n-1, m+n-1; p1 >= 0 || p2 >= 0; tail-- {
//	var cur int
//	if p1 == -1 {
//		cur = nums2[p2]
//		p2--
//	} else if p2 == -1 {
//		cur = nums1[p1]
//		p1--
//	} else if nums1[p1] > nums2[p2] {
//		cur = nums1[p1]
//		p1--
//	} else {
//		cur = nums2[p2]
//		p2--
//	}
//	nums1[tail] = cur
//}
//return nums1

//}

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	arr1 := nums1[:m]
	arr2 := nums2[:n]
	fmt.Println("nums1:", arr1)
	fmt.Println("nums2:", arr2)
	arr := append(arr1, arr2...)
	fmt.Println("arr:", arr)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
