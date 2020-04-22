package _6_minWindos

import "testing"

func TestMinWindow(t *testing.T) {
	t.Log(MinWindow("ADOBECODEBANC", "ABC"))
	t.Log(MinWindow("a", "aa"))
	t.Log(MinWindow2("ADOBECODEBANC", "ABC"))
	t.Log(MinWindow2("a", "aa"))
}
