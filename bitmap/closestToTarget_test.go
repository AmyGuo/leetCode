package bitmap

import "fmt"

//1521. 找到最接近目标值的函数值
//
//func (arr,l,r)  {
//	if (r<l){
//		return -1000000000
//	}
//	ans = arr[l]
//	for (i=l+1；i<r;i++){
//		ans = ans & arr[i]
//	}
//	return ans
//}
//
//Winston 构造了一个如上所示的函数 func 。他有一个整数数组 arr 和一个整数 target ，他想找到让 |func(arr, l, r) - target| 最小的 l 和 r 。
//
//请你返回 |func(arr, l, r) - target| 的最小值。
//
//请注意， func 的输入参数 l 和 r 需要满足 0 <= l, r < arr.length 。
//
//
//
//示例 1：
//输入：arr = [9,12,3,7,15], target = 5
//输出：2
//解释：所有可能的 [l,r] 数对包括 [[0,0],[1,1],[2,2],[3,3],[4,4],[0,1],[1,2],[2,3],[3,4],[0,2],[1,3],[2,4],[0,3],[1,4],[0,4]]， Winston 得到的相应结果为 [9,12,3,7,15,8,0,3,7,0,0,3,0,0,0] 。最接近 5 的值是 7 和 3，所以最小差值为 2 。
//
//
//示例 2：
//输入：arr = [1000000,1000000,1000000], target = 1
//输出：999999
//解释：Winston 输入函数的所有可能 [l,r] 数对得到的函数值都为 1000000 ，所以最小差值为 999999 。
//
//
//示例 3：
//输入：arr = [1,2,4,8,16], target = 0
//输出：0
//
//提示：
//1 <= arr.length <= 10^5
//1 <= arr[i] <= 10^6
//0 <= target <= 10^7

/*
题目解析：
func函数实际是求arr[l]到arr[r]的按位与之和，即 arr[l]&arr[l+1]&..&arr[r-1]&arr[r]
a&b的值不会大于a,也不会大于b，即按位与之和是随着l的减小而单调递减的
按位与之和最多只有20中不同的值
*/

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// 分析:
// 1. 两数相与，只会越来越小;  与得越多，就会越来越小；
// 2. 无序数组 中和目标值与后，最接近目标值的数， 其向两端与更多的数，也会越来越小， 最终会离目标值越来越远；
func closestToTarget(arr []int, target int) int {

	// 1. 找到最接近目标值的数
	midIndex := 0
	mid := target
	dist := abs(arr[0] - target)
	for i := 0; i < len(arr); i++ {
		t := target - arr[i]&target
		if t < mid {
			mid, midIndex = t, i
		}

		// 计算单一值和目标值最小距离
		d := abs(arr[i] - target)
		if d < dist {
			dist = d
			// fmt.Println("i:", i, ",dist:", dist)
		}
	}
	//
	// 判断如果有距离为0直接返回
	if dist == 0 {
		return 0

	}
	// fmt.Println("mid:", midIndex)

	// 向前与，寻找距离更小的结果
	val := arr[midIndex]
	t, tt := val, val
	for i := midIndex - 1; i >= 0; i-- {
		val &= arr[i]
		// fmt.Println("start:", i,"val:",val, ",dist:", dist)
		t = abs(val - target)
		if t > tt {
			// 如果距离变大了，就不用再往前查找了
			break
		}
		tt = t
		if t < dist {
			dist = t
		}
		// fmt.Println("start:", i, ",dist:", dist)
	}

	// 向后与，寻找距离更小的结果
	val = arr[midIndex]
	t, tt = val, val
	for i := midIndex + 1; i < len(arr); i++ {
		val &= arr[i]
		// fmt.Println("end:", i,"val:",val, ",dist:", dist)
		t = abs(val - target)
		if t > tt {
			// 如果距离变大了，就不用再往前查找了
			break
		}
		tt = t
		if t < dist {
			dist = t
		}
		// fmt.Println("end:", i, ",dist:", dist)
	}
	return dist
}

func Example_closestToTarget() {
	fmt.Println(closestToTarget([]int{9, 12, 3, 7, 15}, 5))
	fmt.Println(closestToTarget([]int{1000000, 1000000, 1000000}, 1))
	fmt.Println(closestToTarget([]int{1, 2, 4, 8, 16}, 0))
	//Output:
	//2
	//999999
	//0
}
