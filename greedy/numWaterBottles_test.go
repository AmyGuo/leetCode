package greedy

import "fmt"

/*
1518. 换酒问题
小区便利店正在促销，用 numExchange 个空酒瓶可以兑换一瓶新酒。你购入了 numBottles 瓶酒。
如果喝掉了酒瓶中的酒，那么酒瓶就会变成空的。
请你计算 最多 能喝到多少瓶酒。

示例 1：
输入：numBottles = 9, numExchange = 3
输出：13
解释：你可以用 3 个空酒瓶兑换 1 瓶酒。
所以最多能喝到 9 + 3 + 1 = 13 瓶酒。

示例 2：
输入：numBottles = 15, numExchange = 4
输出：19
解释：你可以用 4 个空酒瓶兑换 1 瓶酒。
所以最多能喝到 15 + 3 + 1 = 19 瓶酒。

示例 3：
输入：numBottles = 5, numExchange = 5
输出：6

示例 4：
输入：numBottles = 2, numExchange = 3
输出：2

提示：

1 <= numBottles <= 100
2 <= numExchange <= 100
*/

func numWaterBottles(numBottles int, numExchange int) int {
	bottle, res := numBottles, numBottles

	for bottle >= numExchange {
		bottle -= numExchange
		res++
		bottle++
	}
	return res
}

func numWaterBottles2(numBottles int, numExchange int) int {
	//一共喝的瓶数
	m := numBottles
	//这里循环一次代表喝了一瓶
	//这里的numBottles代表的是空瓶子，
	for numBottles >= numExchange {
		//喝一次m加一
		m++
		//喝了一瓶代表空瓶数减去numExchange，并且这个时候因为喝了一瓶所以会增加一个空瓶子所以还需要给空瓶加一
		numBottles = numBottles - numExchange + 1
	}
	return m
}

func ExampleNumWaterBottles() {
	fmt.Println(numWaterBottles(9, 3))
	fmt.Println(numWaterBottles(15, 4))
	fmt.Println(numWaterBottles(5, 5))
	fmt.Println(numWaterBottles(2, 3))
	//Output:
	//13
	//19
	//6
	//2
}
