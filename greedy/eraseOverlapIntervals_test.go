package greedy

import (
	"fmt"
)

/*
435. 无重叠区间
给定一个区间的集合，找到需要移除区间的最小数量，使剩余区间互不重叠。

注意:

可以认为区间的终点总是大于它的起点。
区间 [1,2] 和 [2,3] 的边界相互“接触”，但没有相互重叠。
示例 1:

输入: [ [1,2], [2,3], [3,4], [1,3] ]

输出: 1

解释: 移除 [1,3] 后，剩下的区间没有重叠。
示例 2:

输入: [ [1,2], [1,2], [1,2] ]

输出: 2

解释: 你需要移除两个 [1,2] 来使剩下的区间没有重叠。
示例 3:

输入: [ [1,2], [2,3] ]

输出: 0

解释: 你不需要移除任何区间，因为它们已经是无重叠的了。
*/

//先找出最多不重叠的数量，用数组长度减去不重叠的数量，就是需要移除的数量
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	for i := 0; i < len(intervals); i++ {
		for j := i; j < len(intervals); j++ {
			if intervals[i][1] > intervals[j][1] {
				intervals[i], intervals[j] = intervals[j], intervals[i]
			}
		}
	}
	count := 1 //不重叠数量,至少有1个不重叠
	x_end := intervals[0][1]
	for i := 0; i < len(intervals); i++ {
		if intervals[i][0] >= x_end {
			count++
			x_end = intervals[i][1]
		}
	}
	return len(intervals) - count
}

func ExampleEraseOverlapIntervals() {
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}))
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {1, 2}, {1, 2}}))
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {2, 3}}))
	fmt.Println(eraseOverlapIntervals([][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}}))
	//Output:
	//1
	//2
	//0
	//2
}
