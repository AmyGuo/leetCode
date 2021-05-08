package _00_numIslands

/*
200. 岛屿数量
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。

示例 1：

输入：grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
输出：1
示例 2：

输入：grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
输出：3


提示：

m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] 的值为 '0' 或 '1'
*/

//深度优先遍历
func numIslands(grid [][]byte) int {
	num := 0
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < len(grid) && i >= 0 && j < len(grid[0]) && j >= 0 && grid[i][j] == '1' {
			grid[i][j] = '2' //将遍历过的格子标记为2
			dfs(i-1, j)
			dfs(i+1, j)
			dfs(i, j-1)
			dfs(i, j+1)
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				dfs(i, j)
				num++
			}
		}
	}
	return num
}

//广度优先遍历
func numIslands2(grid [][]byte) int {
	num := 0
	var bfs func(i, j int)
	bfs = func(i, j int) {
		queue := [][]int{{i, j}}
		grid[i][j] = '2' //进行标记
		for len(queue) > 0 {
			cur := queue[0]
			aroundData := [][]int{{cur[0] - 1, cur[1]}, {cur[0] + 1, cur[1]}, {cur[0], cur[1] - 1}, {cur[0], cur[1] + 1}}
			for _, idx := range aroundData {
				if idx[0] >= 0 && idx[1] >= 0 && idx[0] < len(grid) && idx[1] < len(grid[0]) && grid[idx[0]][idx[1]] == '1' {
					queue = append(queue, []int{idx[0], idx[1]})
					grid[idx[0]][idx[1]] = '2'
				}
			}
			queue = queue[1:]
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				num++
				bfs(i, j)
			}
		}
	}
	return num
}

//最大岛屿面积
func maxIslandArea(grid [][]byte) int {
	res := 0
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) && grid[i][j] == '1' {
			grid[i][j] = '2'
			return 1 + dfs(i-1, j) + dfs(i+1, j) + dfs(i, j-1) + dfs(i, j+1)
		}
		return 0
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				tmp := dfs(i, j)
				if tmp > res {
					res = tmp
				}
			}
		}
	}
	return res
}

//计算岛屿的周长，网格边长都为1
func maxIslandPerimeter(grid [][]byte) int {
	perimeter := 0
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
			return 1
		}

		if grid[i][j] == '0' {
			return 1
		}

		if grid[i][j] != '1' {
			return 0
		}

		grid[i][j] = '2'
		return dfs(i-1, j) + dfs(i+1, j) + dfs(i, j-1) + dfs(i, j+1)
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				tmp := dfs(i, j)
				if tmp > perimeter {
					perimeter = tmp
				}
			}

		}
	}
	return perimeter
}
