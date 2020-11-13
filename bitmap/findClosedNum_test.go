package bitmap

import "fmt"

//面试题 05.04. 下一个数
//下一个数。给定一个正整数，找出与其二进制表达式中1的个数相同且大小最接近的那两个数（一个略大，一个略小）。
//
//示例1:
//
//输入：num = 2（或者0b10）
//输出：[4, 1] 或者（[0b100, 0b1]）
//示例2:
//
//输入：num = 1
//输出：[2, -1]
//提示:
//
//num的范围在[1, 2147483647]之间；
//如果找不到前一个或者后一个满足条件的正数，那么输出 -1。

//暴力法：
func findClosedNumbers1(num int) []int {
	count := func(n int) int {
		cnt := 0 // 统计n包含1的个数
		for n != 0 {
			cnt++
			n &= n - 1
		}
		return cnt
	}
	tar := count(num)
	max := num + 1
	for {
		if tar == count(max) {
			break // 找到最大值
		}
		max++
	}
	min := num - 1
	for min > 0 {
		if tar == count(min) {
			break // 找到最小值
		}
		min--
	}
	if min == 0 {
		min = -1 // 找不到
	}
	return []int{max, min}
}

func findClosedNumbers(num int) []int {
	bigger, smaller := -1, -1
	bitCount := uint(0)
	for i := uint(0); i != 30 && (bigger == -1 || smaller == -1); i++ {
		if bigger == -1 && (num>>i)&1 == 1 && (num>>(i+1))&1 == 0 {
			bigger = num&^(1<<i) | (1 << (i + 1)) //01--->10
			for j := uint(0); j != i; j++ {       //后面的二进制翻转
				if j < bitCount {
					bigger |= 1 << j //填充够10后面的1的个数
				} else { //其他的都置为0
					bigger &^= 1 << j
				}
			}
		}
		if smaller == -1 && (num>>i)&1 == 0 && (num>>(i+1))&1 == 1 {
			smaller = num&^(1<<(i+1)) | (1 << i)
			for j := uint(0); j != i; j++ {
				if j < i-bitCount {
					smaller &^= 1 << j
				} else {
					smaller |= 1 << j
				}
			}
		}
		if num>>i&1 == 1 { //计算01或者10后面的1的个数
			bitCount++
		}
	}
	return []int{bigger, smaller}
}

//最接近的较小值:二进制表示中从右向左找前面有0的第一个1,"10"中的1,0换位,剩下右边的二进制位前后翻转
//最接近的较大值:二进制表示中从右向左找前面有1的第一个0,"01"中的1,0换位,剩下右边的二进制位前后翻转

//
//比如: 1010011
//略小: 第2个1后面是0则后移一位: 1000, 后面的两个1挨着当前1: 1110（计算方式: 当前1后面的空位全是置为1然后与其右移后面1个数位数做异或, 111^(111>>2)）, 即1001110
//比如: 10110
//略大: 第2个1前面是0则左移一位: 1000, 后面1个1靠右: 1001 (计算方式：后面1的个数填充做或运算, 1<<1 - 1), 即11001

func findClosedNumbers3(num int) []int {
	res := []int{-1, -1}
	flag01, flag10 := true, true
	for i := uint(0); i < 30; i++ {
		curBit, nextBit := 1<<i, 1<<(i+1)
		// 找小值 找01
		if curBit&num == 0 && nextBit&num != 0 && flag01 {
			flag01 = false
			res[1] = num + curBit - nextBit
			// 处理右边的1
			left, j := curBit, uint(0)
			for right := 1; right < left; right <<= 1 {
				if res[1]&right == 0 {
					res[1] = res[1] + right - (1 << j)
					j++
				}
			}
		}
		//  找10
		if curBit&num != 0 && nextBit&num == 0 && flag10 {
			flag10 = false
			res[0] = num - curBit + nextBit
			left, j := curBit, uint(0)
			for right := 1; right < left; right <<= 1 {
				if res[0]&right != 0 {
					res[0] = res[0] - right + (1 << j)
					j++
				}
			}
		}
	}

	return res
}

func Example_findClosedNumbers() {
	fmt.Println(findClosedNumbers(2))
	fmt.Println(findClosedNumbers(1))
	fmt.Println(findClosedNumbers(83))
	fmt.Println(findClosedNumbers(78))
	//Output:
	//[4 1]
	//[2 -1]
	//[85 78]
}
