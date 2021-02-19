package binarySearch

import "fmt"

/*
1337. 矩阵中战斗力最弱的 K 行
给你一个大小为 m * n 的矩阵 mat，矩阵由若干军人和平民组成，分别用 1 和 0 表示。

请你返回矩阵中战斗力最弱的 k 行的索引，按从最弱到最强排序。

如果第 i 行的军人数量少于第 j 行，或者两行军人数量相同但 i 小于 j，那么我们认为第 i 行的战斗力比第 j 行弱。

军人 总是 排在一行中的靠前位置，也就是说 1 总是出现在 0 之前。



示例 1：

输入：mat =
[[1,1,0,0,0],
 [1,1,1,1,0],
 [1,0,0,0,0],
 [1,1,0,0,0],
 [1,1,1,1,1]],
k = 3
输出：[2,0,3]
解释：
每行中的军人数目：
行 0 -> 2
行 1 -> 4
行 2 -> 1
行 3 -> 2
行 4 -> 5
从最弱到最强对这些行排序后得到 [2,0,3,1,4]
示例 2：

输入：mat =
[[1,0,0,0],
 [1,1,1,1],
 [1,0,0,0],
 [1,0,0,0]],
k = 2
输出：[0,2]
解释：
每行中的军人数目：
行 0 -> 1
行 1 -> 4
行 2 -> 1
行 3 -> 1
从最弱到最强对这些行排序后得到 [0,2,3,1]


提示：

m == mat.length
n == mat[i].length
2 <= n, m <= 100
1 <= k <= m
matrix[i][j] 不是 0 就是 1
*/

func kWeakestRows1(mat [][]int, k int) []int {
	numM := make(map[int][]int)
	for k, v := range mat {
		num := 0
		for _, value := range v {
			if value == 1 {
				num++
			}
		}
		numM[num] = append(numM[num], k)
	}

	ans := make([]int, 0, k)
	for i := 0; i <= len(mat[0]); i++ {
		ans = append(ans, numM[i]...)
		if len(ans) >= k {
			ans = ans[:k]
			break
		}
	}
	return ans
}

func kWeakestRows(mat [][]int, k int) []int {
	ans := make([]int, 0)
	isadd := make([]bool, len(mat))
	for i := 0; i < len(mat[0]); i++ {
		for j := 0; j < len(mat); j++ {
			if !isadd[j] && mat[j][i] == 0 && k != 0 {
				ans = append(ans, j)
				isadd[j] = true
				k--
			}
		}
	}
	if k != 0 {
		j := 0
		for k != 0 {
			if !isadd[j] {
				ans = append(ans, j)
				isadd[j] = true
				k--
			}
			j++
		}
	}
	return ans
}

func ExampleKWeakestRows() {
	fmt.Println(kWeakestRows([][]int{{1, 1, 0, 0, 0}, {1, 1, 1, 1, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 0}, {1, 1, 1, 1, 1}}, 3))
	fmt.Println(kWeakestRows([][]int{{1, 0, 0, 0}, {1, 1, 1, 1}, {1, 0, 0, 0}, {1, 0, 0, 0}}, 2))
	fmt.Println(kWeakestRows([][]int{{1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1}}, 1))
	//Output:
	//[2 0 3]
	//[0 2]
	//[0]
}
