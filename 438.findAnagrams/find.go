package _38_findAnagrams

//给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。
//
//字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。
//
//说明：
//
//字母异位词指字母相同，但排列不同的字符串。
//不考虑答案输出的顺序。
//示例 1:
//
//输入:
//s: "cbaebabacd" p: "abc"
//
//输出:
//[0, 6]
//
//解释:
//起始索引等于 0 的子串是 "cba", 它是 "abc" 的字母异位词。
//起始索引等于 6 的子串是 "bac", 它是 "abc" 的字母异位词。
// 示例 2:
//
//输入:
//s: "abab" p: "ab"
//
//输出:
//[0, 1, 2]
//
//解释:
//起始索引等于 0 的子串是 "ab", 它是 "ab" 的字母异位词。
//起始索引等于 1 的子串是 "ba", 它是 "ab" 的字母异位词。
//起始索引等于 2 的子串是 "ab", 它是 "ab" 的字母异位词。

//自己最开始写的有问题的算法
//func FindAnagrams(s string, p string) []int {
//	ss := []byte(s)
//	pp := []byte(p)
//	ssl := len(ss)
//	ppl := len(pp)
//	ret := make([]int, 0, ssl)
//	psum := sum(pp)
//	for i := 0; i <= ssl-ppl; i++ {
//		if sum(ss[i:i+ppl]) == psum {
//			if psum == 0 && ss[i] != pp[0] {
//				continue
//			}
//			ret = append(ret, i)
//		}
//	}
//	return ret
//}
//
//func sum(sli []byte) uint {
//	ret := uint(0)
//	for i := 0; i < len(sli); i++ {
//		for j := i + 1; j < len(sli); j++ {
//			if sli[i] > sli[j] {
//				ret += uint(sli[i] - sli[j])
//			} else {
//				ret += uint(sli[j] - sli[i])
//			}
//		}
//	}
//	return ret
//}

//双指针思想，类似题76
func FindAnagrams(s string, p string) []int {
	ret := make([]int, 0)
	left, right, match := 0, 0, 0
	needs := make(map[byte]int)
	window := make(map[byte]int)

	for i := 0; i < len(p); i++ {
		needs[p[i]]++
	}

	for right < len(s) {
		if _, ok := needs[s[right]]; ok {
			window[s[right]]++
			if window[s[right]] == needs[s[right]] {
				match++
			}
		}
		right++

		for match == len(needs) {
			if right-left == len(p) {
				ret = append(ret, left)
			}
			if v, ok := needs[s[left]]; ok {
				window[s[left]]--
				if window[s[left]] < v {
					match--
				}
			}
			left++
		}

	}
	return ret
}
