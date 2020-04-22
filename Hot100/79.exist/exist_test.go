package _9_exist

import (
	"fmt"
	"testing"
)

func TestExist(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	//word := "SEE"
	//word := "ABCB"
	fmt.Println(Exist(board, word))

}
