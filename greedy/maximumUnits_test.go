package greedy

import (
	"fmt"
	"sort"
)

/*
1710. 卡车上的最大单元数
请你将一些箱子装在 一辆卡车 上。给你一个二维数组 boxTypes ，其中 boxTypes[i] = [numberOfBoxesi, numberOfUnitsPerBoxi] ：

numberOfBoxesi 是类型 i 的箱子的数量。
numberOfUnitsPerBoxi 是类型 i 每个箱子可以装载的单元数量。
整数 truckSize 表示卡车上可以装载 箱子 的 最大数量 。只要箱子数量不超过 truckSize ，你就可以选择任意箱子装到卡车上。

返回卡车可以装载 单元 的 最大 总数。


示例 1：

输入：boxTypes = [[1,3],[2,2],[3,1]], truckSize = 4
输出：8
解释：箱子的情况如下：
- 1 个第一类的箱子，里面含 3 个单元。
- 2 个第二类的箱子，每个里面含 2 个单元。
- 3 个第三类的箱子，每个里面含 1 个单元。
可以选择第一类和第二类的所有箱子，以及第三类的一个箱子。
单元总数 = (1 * 3) + (2 * 2) + (1 * 1) = 8


示例 2：
输入：boxTypes = [[5,10],[2,5],[4,7],[3,9]], truckSize = 10
输出：91


提示：

1 <= boxTypes.length <= 1000
1 <= numberOfBoxesi, numberOfUnitsPerBoxi <= 1000
1 <= truckSize <= 10^6
*/

func maximumUnits(boxTypes [][]int, truckSize int) int {
	res := 0
	//按照可装载的单元排序，再计算
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})

	for i := 0; i < len(boxTypes); i++ {
		if truckSize >= boxTypes[i][0] {
			res += boxTypes[i][0] * boxTypes[i][1]
			truckSize -= boxTypes[i][0]
		} else {
			res += truckSize * boxTypes[i][1]
			break
		}

	}
	return res
}

func maximumUnits2(boxTypes [][]int, truckSize int) int {
	res := 0
	var count = [1001]int{}
	for i, _ := range boxTypes {
		count[boxTypes[i][1]] += boxTypes[i][0]
	}
	left := truckSize
	for i := 1000; i > 0; i-- {
		if count[i] != 0 {
			if count[i] < left {
				res += count[i] * i
				left -= count[i]
			} else {
				res += i * left
				break
			}
		}
	}
	return res
}

func ExampleMaximumUnits() {
	fmt.Println(maximumUnits([][]int{{1, 3}, {2, 2}, {3, 1}}, 4))
	fmt.Println(maximumUnits([][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}, 10))
	//Output:
	//8
	//91
}
