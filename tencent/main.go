package tencent

import (
	"math"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//1. 两数相加:非空链表，逆序存储，将链表对应的两个数相加，返回链表: 输入 2->4->3 5->6->4 返回 7->0->8
func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	var head *ListNode
	var tail *ListNode //用于记录最末尾的结点的地址，最开始tail是指向头结点的
	carry := 0         //进位
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return head
}

//2.寻找两个正序数组的中位数:给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。要求时间复杂度为O(log (m+n))
//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//
//}

//3. 找出最长回文子串【两种情况，中心为奇数，中心为偶数】
func longestPalindrome(s string) string {
	var helper = func(left, right int) (int, int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		return left + 1, right - 1
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		l, r := helper(i, i)
		if r-l > end-start {
			start, end = l, r
		}
	}

	for i := 0; i < len(s); i++ {
		l, r := helper(i, i+1)
		if r-l > end-start {
			start, end = l, r
		}
	}
	return s[start : end+1]
}

//4.给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。如果反转后整数超过 32 位的有符号整数的范围 [−2^31,  2^31 − 1] ，就返回 0。假设环境不允许存储 64 位整数（有符号或无符号）
func reverse(x int) int {
	ans := 0
	for x != 0 {
		if ans < math.MinInt32/10 || ans > math.MaxInt32/10 { //判断乘10时候会溢出，就把该数和最大数除10做比较
			return 0
		}
		dig := x % 10
		x /= 10
		ans = ans*10 + dig
	}
	return ans
}

//5. 字符串转换整数 (atoi)
func myAtoi(s string) int {
	return convert(cleanString(s))
}

func cleanString(s string) (sign int, abs string) {
	//1. 去除首尾的空格
	s = strings.TrimSpace(s)
	if s == "" {
		return
	}

	//2. 判断第一个字符是 符号，数字，其他
	switch s[0] {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		sign, abs = 1, s
	case '+':
		sign, abs = 1, s[1:]
	case '-':
		sign, abs = -1, s[1:]
	default:
		return
	}

	//3. 去除数字+字母模式下的字母
	for i, v := range abs {
		if v < '0' || v > '9' {
			abs = abs[:i]
			break
		}
	}
	return
}

func convert(sign int, abs string) int {
	res := 0
	for _, v := range abs {
		res = res*10 + int(v-'0')
		switch {
		case sign == 1 && res > math.MaxInt32:
			return math.MaxInt32
		case sign == -1 && sign*res < math.MinInt32:
			return math.MinInt32
		}
	}
	return sign * res
}

func myAtoi2(s string) int {
	//1. 去除空格
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}

	//2. 第一位校验
	sign, abs := 0, ""
	switch {
	case s[0] == '+':
		sign, abs = 1, s[1:]
	case s[0] == '-':
		sign, abs = -1, s[1:]
	case s[0] >= '0' && s[0] <= '9':
		sign, abs = 1, s
	}

	//3. 去除数字之后的字母
	for i, v := range abs {
		if v > '9' || v < '0' {
			abs = abs[:i]
			break
		}
	}

	//4. 字符串转数字，并判断溢出
	res := 0
	for _, v := range abs {
		res = res*10 + int(v-'0')
		switch {
		case sign == 1 && res > math.MaxInt32:
			return math.MaxInt32
		case sign == -1 && res*sign < math.MinInt32:
			return math.MinInt32
		}
	}
	return res * sign
}

//6.

//47. 除自身以外数组的乘积 : 给你一个长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。
//输入: [1,2,3,4]
//输出: [24,12,8,6]
//提示：题目数据保证数组之中任意元素的全部前缀元素和后缀（甚至是整个数组）的乘积都在 32 位整数范围内。
//说明: 请不要使用除法，且在 O(n) 时间复杂度内完成此题。
//进阶：
//你可以在常数空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）

func productExceptSelf(nums []int) []int {
	length := len(nums)
	left, right, ret := make([]int, length), make([]int, length), make([]int, length)

	left[0] = 1
	right[length-1] = 1

	for i := 1; i < length; i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	for i := length - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	for i := range ret {
		ret[i] = left[i] * right[i]
	}
	return ret
}

func productExceptSelf2(nums []int) []int {
	length := len(nums)
	answer := make([]int, length)

	// answer[i] 表示索引 i 左侧所有元素的乘积
	// 因为索引为 '0' 的元素左侧没有元素， 所以 answer[0] = 1
	answer[0] = 1
	for i := 1; i < length; i++ {
		answer[i] = nums[i-1] * answer[i-1]
	}

	// R 为右侧所有元素的乘积
	// 刚开始右边没有元素，所以 R = 1
	R := 1
	for i := length - 1; i >= 0; i-- {
		// 对于索引 i，左边的乘积为 answer[i]，右边的乘积为 R
		answer[i] = answer[i] * R
		// R 需要包含右边所有的乘积，所以计算下一个结果时需要将当前值乘到 R 上
		R *= nums[i]
	}
	return answer
}

//48.Nim 游戏: 桌子上有一堆石头。 你们轮流进行自己的回合，你作为先手。 每一回合，轮到的人拿掉 1 - 3 块石头。 拿掉最后一块石头的人就是获胜者。
//假设你们每一步都是最优解。请编写一个函数，来判断你是否可以在给定石头数量为 n 的情况下赢得游戏。如果可以赢，返回 true；否则，返回 false 。
//输入：n = 4
//输出：false
//解释：如果堆中有 4 块石头，那么你永远不会赢得比赛；
//     因为无论你拿走 1 块、2 块 还是 3 块石头，最后一块石头总是会被你的朋友拿走。

func canWinNim(n int) bool {
	return (n%4 != 0)
}

//49.反转字符串 :编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。 不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
//输入：["h","e","l","l","o"]
//输出：["o","l","l","e","h"]
func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	return
}

//50.反转字符串中的单词 III： 给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
//输入："Let's take LeetCode contest"
//输出："s'teL ekat edoCteeL tsetnoc"
func reverseWords(s string) string {
	sBytes := []byte(s)

	//contest LeetCode take Let's
	//left, right := 0, len(s)-1
	//for left < right {
	//	sBytes[left], sBytes[right] = sBytes[right], sBytes[left]
	//	left++
	//	right--
	//}

	for i := 0; i < len(sBytes); {
		start := i
		for i < len(sBytes) && sBytes[i] != ' ' {
			i++
		}
		left, right := start, i-1
		for left < right {
			sBytes[left], sBytes[right] = sBytes[right], sBytes[left]
			left++
			right--
		}

		for i < len(sBytes) && sBytes[i] == ' ' {
			i++
		}
	}
	return string(sBytes)
}
