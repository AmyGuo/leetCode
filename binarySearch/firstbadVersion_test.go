package binarySearch

import "fmt"

/*
278. 第一个错误的版本
你是产品经理，目前正在带领一个团队开发新的产品。不幸的是，你的产品的最新版本没有通过质量检测。由于每个版本都是基于之前的版本开发的，所以错误的版本之后的所有版本都是错的。

假设你有 n 个版本 [1, 2, ..., n]，你想找出导致之后所有版本出错的第一个错误的版本。

你可以通过调用 bool isBadVersion(version) 接口来判断版本号 version 是否在单元测试中出错。实现一个函数来查找第一个错误的版本。你应该尽量减少对调用 API 的次数。

示例:

给定 n = 5，并且 version = 4 是第一个错误的版本。

调用 isBadVersion(3) -> false
调用 isBadVersion(5) -> true
调用 isBadVersion(4) -> true

所以，4 是第一个错误的版本。
*/

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func isBadVersion(version int) bool {
	return version == 4
}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

//这是一道二分查找变形，类似查找某个数第一次出现在递增数组中的位置。且初始值是1，和数组索引不同。所以left=1。

//本题需要注意的是：当中间值不是bug版本时，bug版本一定在递增数组的右边，故收缩左边界left=mid+1。其他情况right=mid，因为right=n 所以for循环内条件为left<right而不是<=
func firstBadVersion(n int) int {
	left, right := 1, n
	for left < right {
		mid := (right-left)/2 + left
		if !isBadVersion(mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func ExampleFirstBadVersion() {
	fmt.Println(firstBadVersion(5))
	//Output:
	//4
}
