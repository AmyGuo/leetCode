package _98_rob

import (
	"fmt"
	"testing"
)

func TestRob(t *testing.T) {
	fmt.Println(Rob([]int{1, 2, 3, 1}))
	fmt.Println(Rob([]int{2, 7, 9, 3, 1}))
	fmt.Println(Rob([]int{2, 1, 1, 2}))

	fmt.Println(Rob2([]int{1, 2, 3, 1}))
	fmt.Println(Rob2([]int{2, 7, 9, 3, 1}))
	fmt.Println(Rob2([]int{2, 1, 1, 2}))

	fmt.Println(Rob3([]int{1, 2, 3, 1}))
	fmt.Println(Rob3([]int{2, 7, 9, 3, 1}))
	fmt.Println(Rob3([]int{2, 1, 1, 2}))
}
