package bitmap

import "fmt"

//面试题 16.01. 交换数字
//编写一个函数，不用临时变量，直接交换numbers = [a, b]中a与b的值。
//
//示例：
//
//输入: numbers = [1,2]
//输出: [2,1]
//提示：
//
//numbers.length == 2

//a ^ b ^ b = a
//a ^ b ^ a = b
//
//a ^ a = 0
//0 ^ a = a
func swapNumbers(numbers []int) []int {
	numbers[0] ^= numbers[1]
	numbers[1] ^= numbers[0]
	numbers[0] ^= numbers[1]
	return numbers
}
func swapNumbers1(numbers []int) []int {
	numbers[0] = numbers[0] + numbers[1]
	numbers[1] = numbers[0] - numbers[1]
	numbers[0] = numbers[0] - numbers[1]
	return numbers
}

func swapNumbers2(numbers []int) []int {
	numbers[0] = numbers[0] - numbers[1]
	numbers[1] = numbers[0] + numbers[1]
	numbers[0] = numbers[1] - numbers[0]
	return numbers
}

func Example_swapNumbers() {
	fmt.Println(swapNumbers([]int{1, 2}))
	fmt.Println(swapNumbers1([]int{1, 2}))
	fmt.Println(swapNumbers2([]int{1, 2}))
	//Output:
	//[2 1]
	//[2 1]
	//[2 1]
}
