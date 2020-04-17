package _87_findDuplicate

//给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。
//
//示例 1:
//
//输入: [1,3,4,2,2]
//输出: 2
//示例 2:
//
//输入: [3,1,3,4,2]
//输出: 3
//说明：
//
//不能更改原数组（假设数组是只读的）。
//只能使用额外的 O(1) 的空间。
//时间复杂度小于 O(n2) 。
//数组中只有一个重复的数字，但它可能不止重复出现一次。

//空间复杂度不符合要求，为O(n)
func FindDuplicate(nums []int) int {
	ret := make([]int, len(nums))
	for _, v := range nums {
		if ret[v] > 0 {
			return v
		}
		ret[v]++
	}
	return 0
}

//把nums看成是顺序存储的链表，nums中每个元素的值是下一个链表节点的地址
//那么如果nums有重复值，说明链表存在环，本问题就转化为了找链表中环的入口节点，因此可以用快慢指针解决
func FindDuplicate2(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast { //先找到两个值相同时对应的下标值
		slow, fast = nums[slow], nums[nums[fast]]
	}

	fast = 0
	for slow != fast { //将两个指针一个指在起始点，一个指在相遇点，两个指针同时走，相遇即为环的入口处，即为重复的数字
		slow, fast = nums[slow], nums[fast]
	}
	return slow
}
