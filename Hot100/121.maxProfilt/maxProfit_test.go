package _21_maxProfilt

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	fmt.Println(MaxProfit([]int{7, 6, 4, 3, 1}))
	fmt.Println(MaxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(MaxProfit2([]int{7, 6, 4, 3, 1}))
	fmt.Println(MaxProfit2([]int{7, 1, 5, 3, 6, 4}))
}
