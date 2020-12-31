package greedy

import (
	"container/heap"
	"fmt"
	"sort"
)

/*
1046. 最后一块石头的重量
有一堆石头，每块石头的重量都是正整数。

每一回合，从中选出两块 最重的 石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的可能结果如下：

如果 x == y，那么两块石头都会被完全粉碎；
如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x。
最后，最多只会剩下一块石头。返回此石头的重量。如果没有石头剩下，就返回 0。



示例：

输入：[2,7,4,1,8,1]
输出：1
解释：
先选出 7 和 8，得到 1，所以数组转换为 [2,4,1,1,1]，
再选出 2 和 4，得到 2，所以数组转换为 [2,1,1,1]，
接着是 2 和 1，得到 1，所以数组转换为 [1,1,1]，
最后选出 1 和 1，得到 0，最终数组转换为 [1]，这就是最后剩下那块石头的重量。


提示：

1 <= stones.length <= 30
1 <= stones[i] <= 1000
*/

//最大堆
type stoneHp struct {
	sort.IntSlice
}

func (s stoneHp) Less(i, j int) bool {
	return s.IntSlice[i] > s.IntSlice[j]
}

func (s *stoneHp) Push(v interface{}) {
	s.IntSlice = append(s.IntSlice, v.(int))
}

func (s *stoneHp) Pop() interface{} {
	tp := s.IntSlice
	v := tp[len(tp)-1]
	s.IntSlice = tp[:len(tp)-1]
	return v
}

func (s *stoneHp) push(v int) {
	heap.Push(s, v)
}

func (s *stoneHp) pop() int {
	return heap.Pop(s).(int)
}

func lastStoneWeight(stones []int) int {
	q := &stoneHp{stones}
	heap.Init(q)
	for q.Len() > 1 {
		a, b := q.pop(), q.pop()
		if a > b {
			q.push(a - b)
		}
	}
	if q.Len() > 0 {
		return q.pop()
	}
	return 0
}

func lastStoneWeight2(stones []int) int {
	sub, length := 0, len(stones)
	for sub+1 < length {

		for i := sub; i < sub+2; i++ {
			for j := i + 1; j < length; j++ {
				if stones[i] < stones[j] {
					stones[i], stones[j] = stones[j], stones[i]
				}
			}
		}

		if stones[sub] == 0 {
			return stones[sub]
		} else if stones[sub] == stones[sub+1] {
			sub = sub + 2
		} else {
			stones[sub+1] = stones[sub] - stones[sub+1]
			sub = sub + 1
		}
	}
	if sub >= length {
		return 0
	}
	return stones[sub]
}

func ExampleLastStoneWeight() {
	fmt.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
	fmt.Println(lastStoneWeight2([]int{2, 7, 4, 1, 8, 1}))
	//Output:
	//1
	//1
}
