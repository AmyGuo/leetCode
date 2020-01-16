package _1_mergeTwoLists

import (
	"fmt"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	//输入：1->2->4, 1->3->4
	//输出：1->1->2->3->4->4
	l1 := new(ListNode)
	l1.Val = 1
	l1.Next = &ListNode{Val: 2, Next: &ListNode{Val: 4}}

	l2 := new(ListNode)
	l2.Val = 1
	l2.Next = &ListNode{Val: 3, Next: &ListNode{Val: 4}}

	ret := MergeTwoLists2(l1, l2)
	//ret := MergeTwoLists(l1, l2)
	for ret != nil {
		fmt.Println(ret.Val)
		ret = ret.Next
	}
}
