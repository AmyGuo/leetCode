package _34_isPalindrome

//请判断一个链表是否为回文链表。
//
//示例 1:
//
//输入: 1->2
//输出: false
//示例 2:
//
//输入: 1->2->2->1
//输出: true

type ListNode struct {
	Val  int
	Next *ListNode
}

//将链表中的值放入切片中，这样就可以很好的判断
func IsPalindrome(head *ListNode) bool {
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-1-i] {
			return false
		}
	}
	return true
}

//利用快慢指针，找到链表的中点，然后翻转后半段链表，遍历前半段链表和后半段链表是否一致，最后再将链表的后半段恢复
func IsPalindrome2(head *ListNode) bool {
	if head == nil {
		return true
	}

	first := EndOfFirstHalf(head)
	second := ReverseList(first.Next)

	p1 := head
	p2 := second

	result := true
	for result && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	first.Next = ReverseList(second)
	return result
}

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

func EndOfFirstHalf(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast.Next.Next != nil && slow.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
