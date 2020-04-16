package _39_wordBreak

import "testing"


func TestWordBreak(t *testing.T) {
	t.Log(WordBreak("leetcode",[]string{"leet","code"}))
	t.Log(WordBreak("applepenapple",[]string{"apple", "pen"}))
	t.Log(WordBreak("catsandog",[]string{"cats", "dog", "sand", "and", "cat"}))
}
