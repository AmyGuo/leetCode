package bitmap

import "fmt"

//面试题 17.01. 不用加号的加法
//设计一个函数把两个数字相加。不得使用 + 或者其他算术运算符。
//
//示例:
//
//输入: a = 1, b = 1
//输出: 2

func add(a int, b int) int {
	for b != 0 {
		temp := (a & b) << 1
		a ^= b
		b = temp
	}
	return a
}

func add2(a int, b int) int {
	for b != 0 {
		a, b = a^b, (a&b)<<1
	}
	return a
}

func Example_add() {
	fmt.Println(add(3, 2))
	//Output:
	//3
}
