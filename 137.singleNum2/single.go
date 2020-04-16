package _37_singleNum2

//
//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只出现了一次的元素。
//
//说明：
//
//你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
//
//示例 1:
//
//输入: [2,2,3,2]
//输出: 3
//示例 2:
//
//输入: [0,1,0,1,0,1,99]
//输出: 99

//0 ^ x = x
//x ^ x = 0
//x & (^x) = 0
//x & (^0) = x
//电路逻辑变化：00->01->10->00
func SingleNumber(nums []int) int {
	a, b := 0, 0
	for _, v := range nums {
		a = (a ^ v) & ^b //前为异或 后为取反
		b = (b ^ v) & ^a
	}
	return a
}

//统计所有数字中每个位中1出现的总数，那么对于某个位，1出现的次数一定是3的倍数+1或0，那么对这个数%3得到的结果就是目的数字在该位上的值
func SingleNumber2(nums []int) int {
	var rec [32]int
	for _, n := range nums {
		for i := uint(0); i < 32; i++ {
			rec[i] = (rec[i] + (n>>i)&1) % 3
		}
	}
	ans := 0
	for i := uint(0); i < 32; i++ {
		ans ^= rec[i] << i
	}
	return int(int32(ans))
}

//3(a+b+c)-(a+a+a+b+b+b+c) = 2c
func SingleNumber3(nums []int) int {
	m := make(map[int]int)

	sum := 0
	for _, v := range nums {
		if _, ok := m[v]; !ok {
			sum += v
		}
		m[v]++
	}
	arrSum := 0
	for i, v := range m {
		arrSum += v * i
	}
	return (3*sum - arrSum) / 2
}
