package bitmap

import (
	"fmt"
)

//面试题 16.07. 最大数值
//编写一个方法，找出两个数字a和b中最大的那一个。不得使用if-else或其他比较运算符。
//
//示例：
//
//输入： a = 1, b = 2
//输出： 2

//如果a>b,则ret=a-0
//如果a<b,则ret=a-(a-b)
//即 a-b > 0 ,a > b  经过处理让a-b=0,a-b < 0 , a < b 经过处理让a-b = a-b。 所以这道题变成了a-b的符号问题。
//负数的最高位是1，正数的最高位是0，那就获取最高位然后和a-b本身相与即可得到答案。
func maximum(a int, b int) int {
	ret := int64(a - b)
	ret = int64(a) - ret&(ret>>63)
	return int(ret)
}

func maximum2(a int, b int) int {
	ret := a - b
	ret = a - ret&(ret>>(32<<(^uint(0)>>63)))
	return ret
}

func maximum3(a int, b int) int {
	k := int((int64(a) - int64(b)) >> 63 & 1)
	return a*(1-k) + b*k
}

func Example_maximum() {
	fmt.Println(maximum(1, 2))
	fmt.Println(maximum(1, 4))
	//Output:
	//2
	//4
}
