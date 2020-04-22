package _34_isPalindrome

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	list := new(ListNode)
	list.Val = 1
	//list.Next = &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}
	//list.Next = &ListNode{Val: 2}
	fmt.Println(IsPalindrome(list))
}
