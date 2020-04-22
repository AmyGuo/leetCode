package _6_minWindos

import "math"

//
//给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字母的最小子串。
//
//示例：
//
//输入: S = "ADOBECODEBANC", T = "ABC"
//输出: "BANC"
//说明：
//
//如果 S 中不存这样的子串，则返回空字符串 ""。
//如果 S 中存在这样的子串，我们保证它是唯一的答案。

//使用左右双指针，优先滑动右指针，直到包含了所有的所需字符串，然后开始左指针滑动，直到遇到第一个所需字符串，不符合需求，继续滑动右指针。
func MinWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	var left, right, match int
	var res string
	targetMap, currMap := map[byte]int{}, map[byte]int{}

	// 处理目标字母的数组
	for i := 0; i < len(t); i++ {
		targetMap[t[i]]++
	}

	for right < len(s) {
		// 寻找字符
		if _, ok := targetMap[s[right]]; ok {
			currMap[s[right]]++
			if currMap[s[right]] == targetMap[s[right]] {
				match++
			}
		}

		right++

		// 如果达到 match 的字符个数，处理 left 指针
		for match == len(targetMap) {
			if len(res) > right-left || res == "" {
				res = s[left:right]
			}
			if v, ok := targetMap[s[left]]; ok {
				currMap[s[left]]--
				if currMap[s[left]] < v {
					match--
				}
			}
			left++
		}
	}

	return res
}

//思路一样，不一样的写法
func MinWindow2(s string, t string) string {
	var needs = make(map[rune]int, len(t))
	for _, v := range t {
		needs[v]++
	}

	var window = make(map[rune]int, len(needs))

	left, right, matchCnt, minLen, start := 0, 0, 0, math.MaxInt32, 0
	for right < len(s) {
		char := rune(s[right])
		if _, ok := needs[char]; ok {
			window[char]++
			if window[char] == needs[char] {
				matchCnt++
			}
		}
		right++

		for matchCnt == len(needs) {
			if right-left < minLen {
				minLen = right - left
				start = left
			}
			char2 := rune(s[left])
			if _, ok := needs[char2]; ok {
				window[char2]--
				if window[char2] < needs[char2] {
					matchCnt--
				}
			}
			left++
		}
	}
	if minLen != math.MaxInt32 {
		return s[start : start+minLen]
	}
	return ""
}
