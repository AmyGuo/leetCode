package _61_hammingDistance

//两个整数之间的汉明距离指的是这两个数字对应二进制位不同的位置的数目。
//
//给出两个整数 x 和 y，计算它们之间的汉明距离。
//
//注意：
//0 ≤ x, y < 2^31.
//
//示例:
//
//输入: x = 1, y = 4
//
//输出: 2
//
//解释:
//1   (0 0 0 1)
//4   (0 1 0 0)
//      ↑   ↑
//
//上面的箭头指出了对应二进制位不同的位置。

//两个数相异或后的数字，计算数字中1的个数

func HammingDistance(x int, y int) int {
	z := x ^ y
	count := 0
	for z > 0 {
		if z&1 == 1 {
			count++
		}
		z >>= 1
	}
	return count
}

func HammingDistance2(x int, y int) int {
	z := x ^ y
	count := 0
	for z > 0 {
		z = z & (z - 1)
		count++
	}
	return count
}

func HammingDistance3(x int, y int) int {
	x = x ^ y
	x = (x & 0x55555555) + ((x >> 1) & 0x55555555)  //1
	x = (x & 0x33333333) + ((x >> 2) & 0x33333333)  //2
	x = (x & 0x0f0f0f0f) + ((x >> 4) & 0x0f0f0f0f)  //3
	x = (x & 0x00ff00ff) + ((x >> 8) & 0x00ff00ff)  //4
	x = (x & 0x0000ffff) + ((x >> 16) & 0x0000ffff) //5
	return x
}
