package _37_singleNum2

import "testing"

func TestSingleNumber(t *testing.T) {
	t.Log(SingleNumber([]int{1, 2, 3, 3, 3, 2, 2}))
	t.Log(SingleNumber([]int{0, 1, 0, 1, 0, 1, 99}))
	t.Log(SingleNumber([]int{-2, -2, 1, 1, -3, 1, -3, -3, -4, -2}))

	t.Log(SingleNumber2([]int{1, 2, 3, 3, 3, 2, 2}))
	t.Log(SingleNumber2([]int{0, 1, 0, 1, 0, 1, 99}))
	t.Log(SingleNumber2([]int{-2, -2, 1, 1, -3, 1, -3, -3, -4, -2}))

	t.Log(SingleNumber3([]int{1, 2, 3, 3, 3, 2, 2}))
	t.Log(SingleNumber3([]int{0, 1, 0, 1, 0, 1, 99}))
	t.Log(SingleNumber3([]int{-2, -2, 1, 1, -3, 1, -3, -3, -4, -2}))
}
