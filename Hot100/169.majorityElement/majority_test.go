package _69_majorityElement

import (
	"fmt"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	fmt.Println(MajorityElement([]int{3, 2, 3}))
	fmt.Println(MajorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
	fmt.Println(MajorityElement2([]int{3, 2, 3}))
	fmt.Println(MajorityElement2([]int{2, 2, 1, 1, 1, 2, 2}))
}
