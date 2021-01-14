package dynammicPlanning

import "fmt"

/*
392. 判断子序列
给定字符串 s 和 t ，判断 s 是否为 t 的子序列。

字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。

进阶：

如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？

致谢：

特别感谢 @pbrother 添加此问题并且创建所有测试用例。



示例 1：
输入：s = "abc", t = "ahbgdc"
输出：true

示例 2：
输入：s = "axc", t = "ahbgdc"
输出：false


提示：

0 <= s.length <= 100
0 <= t.length <= 10^4
两个字符串都只由小写字符组成。

*/

func isSubsequence(s string, t string) bool {
	n, m := len(s), len(t)
	f := make([][26]int, m+1)
	for i := 0; i < 26; i++ {
		f[m][i] = m
	}
	for i := m - 1; i >= 0; i-- {
		for j := 0; j < 26; j++ {
			if t[i] == byte(j+'a') {
				f[i][j] = i
			} else {
				f[i][j] = f[i+1][j]
			}
		}
	}
	add := 0
	for i := 0; i < n; i++ {
		if f[add][int(s[i]-'a')] == m {
			return false
		}
		add = f[add][int(s[i]-'a')] + 1
	}
	return true
}

func isSubsequence2(s string, t string) bool {
	m, n := len(s), len(t)
	i, j := 0, 0

	for i < m && j < n {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return m == i
}

/*
class Solution {
   public boolean isSubsequence(String s, String t) {
        int sLen = s.length(), tLen = t.length();
        if (sLen > tLen) return false;
        if (sLen == 0) return true;
        boolean[][] dp = new boolean[sLen + 1][tLen + 1];
        for (int j = 0; j <= tLen; j++) {
            dp[0][j] = true;
        }

        for (int i = 1; i <= sLen; i++) {
            for (int j = 1; j <= tLen; j++) {
                if (s.charAt(i - 1) == t.charAt(j - 1)) {
                    dp[i][j] = dp[i - 1][j - 1];
                } else {
                    dp[i][j] = dp[i][j - 1];
                }
            }
        }
        return dp[sLen][tLen];
    }
}

*/

func ExampleIsSubsequence() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence2("abc", "ahbgdc"))
	fmt.Println(isSubsequence("axc", "ahbgdc"))
	fmt.Println(isSubsequence2("axc", "ahbgdc"))
	//Output:
	//true
	//true
	//false
	//false
}
