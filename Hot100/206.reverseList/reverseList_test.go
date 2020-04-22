package _06_reverseList

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {
	list := new(ListNode)
	list.Val = 1
	list.Next = &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}
	rList := ReverseList(list)

	for rList.Next != nil {
		fmt.Printf("%d->", rList.Val)
		rList = rList.Next
	}
	fmt.Printf("\n")
}
