package _60_getIntersectionNode

import (
	"fmt"
	"testing"
)

//listA = [0,9,1,2,4], listB = [3,2,4],
func TestGetIntersectionNode(t *testing.T) {
	listA := new(ListNode)
	listB := new(ListNode)
	com := &ListNode{Val: 2, Next: &ListNode{Val: 4}}
	listA.Val = 0
	listA.Next = &ListNode{Val: 9, Next: &ListNode{Val: 1, Next: com}}
	listB.Val = 3
	listB.Next = com
	list := GetIntersectionNode(listA, listB)
	fmt.Println(list)
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}

}
