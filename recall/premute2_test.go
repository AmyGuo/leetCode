package recall

import (
	"fmt"
	"strconv"
)

/*
60. 排列序列
给出集合 [1,2,3,...,n]，其所有元素共有 n! 种排列。

按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：

"123"
"132"
"213"
"231"
"312"
"321"
给定 n 和 k，返回第 k 个排列。



示例 1：

输入：n = 3, k = 3
输出："213"
示例 2：

输入：n = 4, k = 9
输出："2314"
示例 3：

输入：n = 3, k = 1
输出："123"


提示：

1 <= n <= 9
1 <= k <= n!
*/

func getPermutation(n int, k int) string {
	res := ""
	index := 0

	var dfs func(used map[int]bool, path string)

	dfs = func(used map[int]bool, path string) {
		if index == k {
			return
		}

		if len(path) == n {
			index++
			res = path
			path = ""
			return
		}

		for i := 1; i <= n; i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path += strconv.Itoa(i)
			dfs(used, path)
			temp := []byte(path)
			path = string(temp[:len(temp)-1])
			used[i] = false
		}
	}

	m := make(map[int]bool)
	dfs(m, "")
	return res
}

func Example_getPermutation() {
	fmt.Println(getPermutation(3, 3))
	fmt.Println(getPermutation(4, 9))
	fmt.Println(getPermutation(3, 1))
	//Output:
	//213
	//2314
	//123
}
