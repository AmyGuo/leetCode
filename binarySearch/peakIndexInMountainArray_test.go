package binarySearch

import "fmt"

/*
852. 山脉数组的峰顶索引
符合下列属性的数组 arr 称为 山脉数组 ：
arr.length >= 3
存在 i（0 < i < arr.length - 1）使得：
arr[0] < arr[1] < ... arr[i-1] < arr[i]
arr[i] > arr[i+1] > ... > arr[arr.length - 1]
给你由整数组成的山脉数组 arr ，返回任何满足 arr[0] < arr[1] < ... arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1] 的下标 i 。



示例 1：

输入：arr = [0,1,0]
输出：1
示例 2：

输入：arr = [0,2,1,0]
输出：1
示例 3：

输入：arr = [0,10,5,2]
输出：1
示例 4：

输入：arr = [3,4,5,1]
输出：2
示例 5：

输入：arr = [24,69,100,99,79,78,67,36,26,19]
输出：2


提示：

3 <= arr.length <= 10^4
0 <= arr[i] <= 10^6
题目数据保证 arr 是一个山脉数组


进阶：很容易想到时间复杂度 O(n) 的解决方案，你可以设计一个 O(log(n)) 的解决方案吗？
*/

func peakIndexInMountainArray3(arr []int) int {
	left, right := 1, len(arr)-2
	for left <= right {
		mid := (right-left)>>1 + left
		if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			return mid
		} else if arr[mid] < arr[mid-1] && arr[mid] > arr[mid+1] {
			right = mid - 1
		} else if arr[mid] > arr[mid-1] && arr[mid] < arr[mid+1] {
			left = mid + 1
		}
	}
	return 0
}

func peakIndexInMountainArray(arr []int) int {
	start := 0
	end := len(arr) - 1

	for start < end {
		mid := (end + start) / 2
		if arr[mid] < arr[mid+1] {
			start = mid + 1
		} else {
			end = mid
		}
	}

	return end
}

func peakIndexInMountainArray2(arr []int) int {
	for i := 1; i <= len(arr)-2; i++ {
		if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
			return i
		}
	}
	return 0
}

func ExamplePeakIndexInMountainArray() {
	fmt.Println(peakIndexInMountainArray([]int{0, 1, 0}))
	fmt.Println(peakIndexInMountainArray([]int{0, 2, 1, 0}))
	fmt.Println(peakIndexInMountainArray([]int{0, 10, 5, 2}))
	fmt.Println(peakIndexInMountainArray([]int{3, 4, 5, 1}))
	fmt.Println(peakIndexInMountainArray([]int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}))
	//Output:
	//1
	//1
	//1
	//2
	//2
}
