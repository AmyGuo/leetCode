package _9_exist

//单词搜索
//给定一个二维网格和一个单词，找出该单词是否存在于网格中。
//
//单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
//
//示例:
//
//board =
//[
//['A','B','C','E'],
//['S','F','C','S'],
//['A','D','E','E']
//]
//
//给定 word = "ABCCED", 返回 true.
//给定 word = "SEE", 返回 true.
//给定 word = "ABCB", 返回 false.


//tips:二维平面上使用回溯法
func Exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(board, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, i int, j int, word string, k int) bool {
	if board[i][j] != word[k] {
		return false
	}
	if k == len(word)-1 {
		return true
	}
	temp := board[i][j]
	board[i][j] = byte(' ')
	if 0 <= i-1 && dfs(board, i-1, j, word, k+1) {
		return true
	}
	if i+1 < len(board) && dfs(board, i+1, j, word, k+1) {
		return true
	}
	if 0 <= j-1 && dfs(board, i, j-1, word, k+1) {
		return true
	}
	if j+1 < len(board[0]) && dfs(board, i, j+1, word, k+1) {
		return true
	}
	board[i][j] = temp
	return false
}



//不同的写法，更加简便一些，消耗时间更短(思想和上面的解法一致)
func Exist2(nums [][]byte, word string) bool {
	for i:=0; i < len(nums); i++ {
		for j:=0; j< len(nums[0]); j++ {
			if helper(nums, i, j , word , 0) {
				return true
			}
		}
	}
	return false
}

func helper(nums [][]byte, i, j int,  word string, index int)bool{
	// 1.递归退出条件，index== 单词长度  true
	if len(word) == index {
		return true
	}
	// 2.递归退出条件， 越界
	if i<0 || i==len(nums) || j<0 || j==len(nums[0]) {
		return false
	}
	// 3.递归退出条件, word当前值与nums中不匹配
	if nums[i][j] != word[index] {
		return false
	}
	// 进入递归逻辑
	curr := nums[i][j]
	nums[i][j] = byte(' ') // 防止重复
	ret := helper(nums, i, j+1, word, index+1) ||
		helper(nums, i, j-1, word, index+1) ||
		helper(nums, i+1, j, word, index+1) ||
		helper(nums, i-1, j, word, index+1)
	nums[i][j] = curr // 回溯
	return ret
}