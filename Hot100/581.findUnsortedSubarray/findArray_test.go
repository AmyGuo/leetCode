package _81_findUnsortedSubarray

import (
	"fmt"
	"testing"
)

func TestFindUnsortedSubarray(t *testing.T) {
	fmt.Println(FindUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
	fmt.Println(FindUnsortedSubarray([]int{1, 2, 3, 4}))
	fmt.Println(FindUnsortedSubarray([]int{2, 1}))
	fmt.Println(FindUnsortedSubarray([]int{1, 2, 3, 5, 4}))
	fmt.Println(FindUnsortedSubarray([]int{1, 3, 2, 4, 5}))
	fmt.Println("======")
	fmt.Println(FindUnsortedSubarray3([]int{2, 6, 4, 8, 10, 9, 15}))
	fmt.Println(FindUnsortedSubarray3([]int{1, 2, 3, 4}))
	fmt.Println(FindUnsortedSubarray3([]int{2, 1}))
	fmt.Println(FindUnsortedSubarray3([]int{1, 2, 3, 5, 4}))
	fmt.Println(FindUnsortedSubarray3([]int{1, 3, 2, 4, 5}))
}
