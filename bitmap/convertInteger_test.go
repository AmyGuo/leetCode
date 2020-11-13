package bitmap

import "fmt"

//面试题 05.06. 整数转换
//整数转换。编写一个函数，确定需要改变几个位才能将整数A转成整数B。
//
//示例1:
//
//输入：A = 29 （或者0b 11101）, B = 15（或者0b 01111）
//输出：2
//示例2:
//
//输入：A = 1，B = 2
//输出：2
//提示:
//
//A，B范围在[-2147483648, 2147483647]之间

//异或之后，统计1的个数
func convertInteger(A int, B int) int {
	// 因为右移的话是补符号位，因此出现负数的话就会出错，因此要左移
	count := 0
	tmp := A ^ B
	for tmp != 0 {
		if tmp&0x80000000 == 0x80000000 {
			count++
		}
		tmp <<= 1
	}
	return count
}

func Example_convertInteger() {
	fmt.Println(convertInteger(29, 15))
	fmt.Println(convertInteger(1, 2))
	fmt.Println(convertInteger(826966453, -729934991))
	//Output:
	//2
	//2
	//14
}
