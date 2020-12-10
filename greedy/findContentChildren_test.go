package greedy

import "fmt"

/*
455. 分发饼干
假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干。

对每个孩子 i，都有一个胃口值 g[i]，这是能让孩子们满足胃口的饼干的最小尺寸；并且每块饼干 j，都有一个尺寸 s[j] 。如果 s[j] >= g[i]，我们可以将这个饼干 j 分配给孩子 i ，这个孩子会得到满足。你的目标是尽可能满足越多数量的孩子，并输出这个最大数值。


示例 1:

输入: g = [1,2,3], s = [1,1]
输出: 1
解释:
你有三个孩子和两块小饼干，3个孩子的胃口值分别是：1,2,3。
虽然你有两块小饼干，由于他们的尺寸都是1，你只能让胃口值是1的孩子满足。
所以你应该输出1。
示例 2:

输入: g = [1,2], s = [1,2,3]
输出: 2
解释:
你有两个孩子和三块小饼干，2个孩子的胃口值分别是1,2。
你拥有的饼干数量和尺寸都足以让所有孩子满足。
所以你应该输出2.


提示：

1 <= g.length <= 3 * 10^4
0 <= s.length <= 3 * 10^4
1 <= g[i], s[j] <= 2^31 - 1
*/

//满足尽可能多的孩子的胃口，那么就要从最小的胃口的孩子派发，这样可以得到尽可能多的人
//所以需要对孩子胃口数组和饼干数组同时做排序
func findContentChildren(g []int, s []int) int {
	if len(s) == 0 {
		return 0
	}

	for i := 0; i < len(g); i++ {
		for j := i; j < len(g); j++ {
			if g[i] > g[j] {
				g[i], g[j] = g[j], g[i]
			}
		}
	}

	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}

	i, j := 0, 0

	for i < len(g) && j < len(s) {
		if s[j] >= g[i] {
			i++
		}
		j++
	}
	return i
}

func ExampleFindContentChildren() {
	fmt.Println(findContentChildren([]int{1, 2, 3}, []int{1, 1}))
	fmt.Println(findContentChildren([]int{1, 2}, []int{1, 2, 3}))
	//Output:
	//1
	//2
}
