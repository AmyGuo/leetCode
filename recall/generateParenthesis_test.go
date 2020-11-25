package recall

import (
	"fmt"
	"testing"
)

/*
22. 括号生成
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。


示例：
输入：n = 3
输出：[
       "((()))",
       "(()())",
       "(())()",
       "()(())",
       "()()()"
     ]
*/

func generateParenthesis(n int) []string {
	if n <= 0 {
		return []string{}
	}

	res := []string{}

	var backTrack func(leftIndex, rightIndex int, pth string)
	backTrack = func(leftIndex, rightIndex int, pth string) {
		if len(pth) == n*2 {
			res = append(res, pth)
			pth = ""
			return
		}
		if leftIndex > 0 {
			pth = pth + "("
			backTrack(leftIndex-1, rightIndex, pth)
			pth = pth[:len(pth)-1]
		}

		if rightIndex > leftIndex { //右括号剩余多
			pth = pth + ")"
			backTrack(leftIndex, rightIndex-1, pth)
			pth = pth[:len(pth)-1]
		}
	}

	backTrack(n, n, "")
	return res
}

func generateParenthesis2(n int) []string {
	if n <= 0 {
		return []string{}
	}

	res := []string{}

	var backTrack func(leftIndex, rightIndex int, pth string)
	backTrack = func(leftIndex, rightIndex int, pth string) {
		if len(pth) == n*2 {
			res = append(res, pth)
			pth = ""
			return
		}
		if leftIndex < n {
			pth = pth + "("
			backTrack(leftIndex+1, rightIndex, pth)
			pth = pth[:len(pth)-1]
		}

		if rightIndex < leftIndex { //右括号加的少
			pth = pth + ")"
			backTrack(leftIndex, rightIndex+1, pth)
			pth = pth[:len(pth)-1]
		}
	}

	backTrack(0, 0, "")
	return res
}

func Test_generateParenthesis(t *testing.T) {
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis2(3))
}
