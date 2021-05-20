package tencent

import "math"

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
		if ans < math.MinInt32/10 || ans > math.MinInt32/10 {
			return 0
		}
		dig := x % 10
		x /= 10
		ans = ans*10 + dig
	}
	return ans
}
