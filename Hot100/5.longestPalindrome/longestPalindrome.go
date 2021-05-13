package __longestPalindrome

/*
5. 最长回文子串
给你一个字符串 s，找到 s 中最长的回文子串。



示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：

输入：s = "cbbd"
输出："bb"
示例 3：

输入：s = "a"
输出："a"
示例 4：

输入：s = "ac"
输出："a"
*/

func longestPalindrome(s string) string {
	var helper = func(left, right int) (int, int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		return left + 1, right - 1
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		l, r := helper(i, i)
		if end-start < r-l {
			start, end = l, r
		}
	}
	for i := 0; i < len(s)-1; i++ {
		l, r := helper(i, i+1)
		if end-start < r-l {
			start, end = l, r
		}
	}
	return s[start : end+1]
}
