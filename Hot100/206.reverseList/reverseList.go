package _06_reverseList

//反转一个单链表。
//
//示例:
//
//输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL

type ListNode struct {
	Val  int
	Next *ListNode
}

//迭代
func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}

//递归
//https://zhuanlan.zhihu.com/p/86745433 知乎上的解释，很清晰好理解
func ReverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	res := ReverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil // 破环
	return res
}
