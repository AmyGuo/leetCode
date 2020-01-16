package _1_mergeTwoLists

type ListNode struct {
	Val  int
	Next *ListNode
}

//将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
//示例：
//
//输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4

//递归法：
func MergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = MergeTwoLists2(l1.Next, l2)
		return l1
	} else {
		l2.Next = MergeTwoLists2(l1, l2.Next)
		return l2
	}
}

//迭代法：
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	newList := head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			newList.Next = l1
			l1 = l1.Next
		} else {
			newList.Next = l2
			l2 = l2.Next
		}
		newList = newList.Next
	}
	if l1 == nil {
		newList.Next = l2
	}
	if l2 == nil {
		newList.Next = l1
	}

	return head.Next
}
