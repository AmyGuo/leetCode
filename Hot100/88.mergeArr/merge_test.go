package Hot100

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	fmt.Println(merge([]int{1, 2, 3}, 3, []int{2, 4, 5}, 3))
}
