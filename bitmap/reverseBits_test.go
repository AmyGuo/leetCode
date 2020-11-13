package bitmap

import (
	"fmt"
)

//面试题 05.03. 翻转数位
//给定一个32位整数 num，你可以将一个数位从0变为1。请编写一个程序，找出你能够获得的最长的一串1的长度。
//
//示例 1：
//
//输入: num = 1775(11011101111)
//输出: 8
//示例 2：
//
//输入: num = 7(0111)
//输出: 4

func reverseBits1(num int) int {
	s := []int{}
	count := 0
	str := ""
	if num < 0 {
		str = fmt.Sprintf("%b", int(^uint32(-num)+1))
	} else {
		str = fmt.Sprintf("%b", int(num))
	}
	for _, v := range str {
		if v == '1' {
			count++
		} else {
			s = append(s, count)
			count = 0
		}
	}

	if count > 0 {
		s = append(s, count)
	}
	if len(s) == 1 {
		if s[0] == 32 {
			return s[0]
		} else {
			return s[0] + 1
		}
	}

	max := 0
	for i := 1; i < len(s); i++ {
		if s[i-1]+s[i] > max {
			max = s[i-1] + s[i]
		}
	}
	return max + 1
}

func reverseBits(num int) int {
	reverseMax, curMax, max := 0, 0, 0
	bit := 1
	for i := 0; i != 32; i++ {
		if num|bit == num { //说明第bit位是1
			curMax++
			reverseMax++
		} else {
			if max < reverseMax {
				max = reverseMax
			}
			reverseMax = curMax + 1
			curMax = 0
		}
		bit <<= 1
	}
	if max < reverseMax {
		max = reverseMax
	}
	return max
}

func Example_reverseBits() {
	fmt.Println(reverseBits(1775))
	fmt.Println(reverseBits(7))
	fmt.Println(reverseBits(-1))
	fmt.Println(reverseBits(-2))
	fmt.Println(reverseBits(-4))
	//Output:
	//8
	//4
	//32
	//32
	//31
}
