package bitmap

import "fmt"

//1525. 字符串的好分割数目
//给你一个字符串 s ，一个分割被称为 「好分割」 当它满足：将 s 分割成 2 个字符串 p 和 q ，它们连接起来等于 s 且 p 和 q 中不同字符的数目相同。
//
//请你返回 s 中好分割的数目。

//示例 1：
//输入：s = "aacaba"
//输出：2
//解释：总共有 5 种分割字符串 "aacaba" 的方法，其中 2 种是好分割。
//("a", "acaba") 左边字符串和右边字符串分别包含 1 个和 3 个不同的字符。
//("aa", "caba") 左边字符串和右边字符串分别包含 1 个和 3 个不同的字符。
//("aac", "aba") 左边字符串和右边字符串分别包含 2 个和 2 个不同的字符。这是一个好分割。
//("aaca", "ba") 左边字符串和右边字符串分别包含 2 个和 2 个不同的字符。这是一个好分割。
//("aacab", "a") 左边字符串和右边字符串分别包含 3 个和 1 个不同的字符。

//示例 2：
//
//输入：s = "abcd"
//输出：1
//解释：好分割为将字符串分割成 ("ab", "cd") 。

//示例 3：
//
//输入：s = "aaaaa"
//输出：4
//解释：所有分割都是好分割。

//示例 4：
//
//输入：s = "acbadbaada"
//输出：2
//
//
//提示：
//s 只包含小写英文字母。
//1 <= s.length <= 10^5

//这种解法没有想明白，先记录在这里吧
func numSplits1(s string) int {
	lmap := make(map[rune]int)
	rmap := make(map[rune]int)
	for _, c := range s {
		rmap[c] = rmap[c] + 1
	}
	res := 0
	for _, c := range s {
		lmap[c] = lmap[c] + 1
		rmap[c] = rmap[c] - 1
		if rmap[c] == 0 {
			delete(rmap, c)
		}
		if len(lmap) == len(rmap) {
			res++
		}
	}
	return res
}

/*
官方解法：
动态规划解法:
要使得左右两部分不同字符的数量相同，首先要统计一侧的不同字符数量。
注意到当我们已经统计出字符串的前 i−1 个字符中的不同字符数量时，我们可以利用该信息算出字符串前 i 个字符的不同字符数量。
记字符串的前 i 个字符（为了方便处理边界情况，下标从 1开始编号）中的不同字符数量为 f[i]，当第 i 个字符在字符串前 i−1 个字符中出现过时，f[i]=f[i-1]，否则 f[i]=f[i-1]+1。
特别地，f[0]=0
为了判断第 i 个字符是否在字符串前 i-1个字符中出现过，我们可以用一个布尔类型的数组 rec 进行标记，令 rec[i] 表示字符 i 是否出现过。这样我们只需要在统计的同时，不断更新这个 rec 数组，
即可实现每次 O(1) 地查询字符串中第 i 个字符是否出现过。
实现了统计一侧的不同字符数量的功能，我们可以如法炮制，只需要将动态规划的计算顺序进行反向，就能统计出另一侧的不同字符数量。
最后我们枚举每一个位置，判断其两侧的字符数量是否相同，即可知道这个位置的分割是不是一个「好分割」。
*/
func numSplits2(s string) int {
	left := make([]int, len(s))
	record := make([]bool, 26)
	for i := range s {
		if record[int(s[i]-'a')] {
			left[i] = left[i-1]
		} else {
			record[int(s[i]-'a')] = true
			if i-1 >= 0 {
				left[i] = left[i-1] + 1
			} else {
				left[i] = 1
			}
		}
	}

	result := 0
	record = make([]bool, 26)
	right := 0
	for i := len(s) - 1; i > 0; i-- {
		if !record[int(s[i]-'a')] {
			record[int(s[i]-'a')] = true
			right++
		}
		if left[i-1] == right {
			result++
		}
	}
	return result
}

func numSplits(s string) int {
	left := make([]int, len(s))
	leftRecord := make(map[byte]bool)
	for i := range s {
		leftRecord[s[i]] = true
		left[i] = len(leftRecord)
	}

	result := 0
	rightRecord := make(map[byte]bool)
	for i := len(s) - 1; i > 0; i-- {
		rightRecord[s[i]] = true
		if left[i-1] == len(rightRecord) {
			result++
		}
	}
	return result
}

func Example_NumSplits() {
	fmt.Println(numSplits("aacaba"))
	fmt.Println(numSplits("abcd"))
	fmt.Println(numSplits("aaaaa"))
	fmt.Println(numSplits("acbadbaada"))
	//Output:
	//2
	//1
	//4
	//2
}
