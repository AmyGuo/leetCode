package _41_hasCycle

import (
	"fmt"
	"testing"
)

func TestHasCycle(t *testing.T) {
	node1 := new(ListNode)
	node2 := new(ListNode)
	node3 := new(ListNode)
	node4 := new(ListNode)
	node1.Val = 1
	node2.Val = 2
	node3.Val = 3
	node4.Val = 4
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	//node4.Next = node1
	fmt.Println(HasCycle(node1))
	fmt.Println(HasCycle2(node1))
}
