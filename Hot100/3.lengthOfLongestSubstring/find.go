package __lengthOfLongestSubstring

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
//示例 1:
//
//输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//
//输入: "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

//76, 438题 类似

//滑动窗口抽象思想：
//int left = 0, right = 0;
//
//while (right < s.size()) {
//	window.add(s[right]);
//	right++;
//
//	while (valid) {
//		window.remove(s[left]);
//			left++;
//		}
//}

func LengthOfLongestSubstring(s string) int {
	left, right, res := 0, 0, 0
	window := make(map[byte]int)

	for right < len(s) {
		c1 := s[right]
		window[c1]++
		right++

		for window[c1] > 1 {
			window[s[left]]--
			left++
		}
		if res < right-left {
			res = right - left
		}
	}
	return res
}
