package _39_wordBreak

//给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
//
//说明：
//
//拆分时可以重复使用字典中的单词。
//你可以假设字典中没有重复的单词。
//示例 1：
//
//输入: s = "leetcode", wordDict = ["leet", "code"]
//输出: true
//解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
//示例 2：
//
//输入: s = "applepenapple", wordDict = ["apple", "pen"]
//输出: true
//解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
//     注意你可以重复使用字典中的单词。
//示例 3：
//
//输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
//输出: false


func WordBreak(s string, wordDict []string) bool {
	return true
}

func wordBreak(s string, wordDict []string) bool {
	return bword([]byte(s),wordDict)
}

func bword(s []byte, keys []string) bool {
	dp := make([]bool, len(s))
	for i := 0; i < len(s); i++ {
		for _, key := range keys {
			if i < len(key) - 1 {
				continue
			}
			if (i == len(key) - 1 || dp[i-len(key)]) && string(s[i-len(key)+1:i+1]) == key {
				dp[i] = true
			}
		}
	}

	return dp[len(s)-1]
}
//
//重点思考状态转移方程：
//i为字符串第i位， key为其中一个关键字
//f(i) = f(i - len(key)) 且 key == s[i-len(key)+1:i+1]
//
//用一个与s等长的数组dp保存状态即可，如果f(i)成立dp[i] == true

