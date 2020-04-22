package _60_singleNum3

//给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。
//
//示例 :
//
//输入: [1,2,1,3,2,5]
//输出: [3,5]
//注意：
//
//结果输出的顺序并不重要，对于上面的例子， [5, 3] 也是正确答案。
//你的算法应该具有线性时间复杂度。你能否仅使用常数空间复杂度来实现？
//

func SingleNumber(nums []int) []int {
	temp := 0
	for _, v := range nums {
		temp ^= v
	}
	for _, v := range nums {
		for _, n := range nums {
			if v^n == temp {
				return []int{v, n}
			}
		}
	}
	return nil
}

func SingleNumber2(nums []int) []int {
	bitmask := 0
	for _, v := range nums {
		bitmask ^= v
	}
	diff := bitmask & (-bitmask) //保留bitmask最右边的1，这个1要么来自x，要么来自y,（即使用bitmask来分离x和y）
	ans := make([]int, 2)
	//通过与这个mask进行与操作，如果为0的分为一个数组，为1的分为另一个数组。这样就把问题降低成了：“有一个数组每个数字都出现两次，有一个数字只出现了一次，求出该数字”。对这两个子问题分别进行全异或就可以得到两个解。也就是最终的数组了。
	for _, num := range nums {
		if (num & diff) == 0 {
			ans[0] ^= num
		} else {
			ans[1] ^= num
		}

	}
	return ans
}
