package _87_findDuplicate

import "testing"

func TestFindDuplicate(t *testing.T) {
	t.Log(FindDuplicate([]int{1, 3, 4, 2, 2}))
	t.Log(FindDuplicate([]int{3, 1, 3, 4, 2}))

	t.Log(FindDuplicate2([]int{1, 3, 4, 2, 2}))
	t.Log(FindDuplicate2([]int{3, 1, 3, 4, 2}))
}
