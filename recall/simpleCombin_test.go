package recall

import (
	"fmt"
)

/*
输入两个数字 n, k，算法输出 [1..n] 中 k 个数字的所有组合。
*/

func combine(n, k int) [][]int {
	res := make([][]int, 0)

	var dfs func(index int, path []int)

	dfs = func(index int, path []int) {
		if len(path) == k {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i := index; i <= n; i++ {
			path = append(path, i)
			dfs(i+1, path)
			path = path[:len(path)-1]
		}
	}

	dfs(1, []int{})
	return res
}

func Example_combine() {
	fmt.Println(combine(4, 2))
	//Output:
	//[[1 2] [1 3] [1 4] [2 3] [2 4] [3 4]]
}
