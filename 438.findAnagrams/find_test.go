package _38_findAnagrams

import (
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	t.Log(FindAnagrams("abab", "ab"))
	t.Log(FindAnagrams("cbaebabacd", "abc"))
	t.Log(FindAnagrams("aa", "bb"))
	t.Log(FindAnagrams("bpaa", "aa"))
	t.Log(FindAnagrams("abaacbabc", "abc"))
}
