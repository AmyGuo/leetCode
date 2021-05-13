package _47_countSubStrings

//647. 回文子串
//给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。
//
//具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
//
//
//
//示例 1：
//
//输入："abc"
//输出：3
//解释：三个回文子串: "a", "b", "c"
//示例 2：
//
//输入："aaa"
//输出：6
//解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"

func countSubstrings(s string) int {
	var helper func(left, right int) int
	helper = func(left, right int) int {
		count := 0
		for left >= 0 && right <= len(s)-1 && s[left] == s[right] {
			left--
			right++
			count++
		}
		return count
	}
	res := 0
	for i := 0; i < len(s); i++ {
		res += helper(i, i)
	}
	for i := 0; i < len(s)-1; i++ {
		res += helper(i, i+1)
	}
	return res
}
