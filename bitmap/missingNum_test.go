package bitmap

import (
	"fmt"
	"testing"
)

//面试题 17.04. 消失的数字
//数组nums包含从0到n的所有整数，但其中缺了一个。请编写代码找出那个缺失的整数。你有办法在O(n)时间内完成吗？
//
//注意：本题相对书上原题稍作改动
//
//示例 1：
//
//输入：[3,0,1]
//输出：2
//
//
//示例 2：
//
//输入：[9,6,4,2,3,5,7,0,1]
//输出：8

//A^A=0 0^A=A
//0-n的整数对应的下标是0-(n-1),如果全部不缺，要让所有数全部出现两次，就加上最后异或数组长度n
func missingNumber(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result ^= i
		result ^= nums[i]
	}
	result ^= len(nums)
	return result
}

func missingNumber2(nums []int) int {
	sum, all := 0, 0
	for i, num := range nums {
		sum += num
		all += i + 1
	}
	return all - sum
}

func Example_missingNumber() {
	fmt.Println(missingNumber([]int{3, 0, 1}))
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
	//Output:
	//2
	//8
}

func Benchmark_missingNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ret := missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})
		if ret != 8 {
			panic("err!")
		}
	}
}

func Benchmark_missingNumber2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ret := missingNumber2([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})
		if ret != 8 {
			panic("err!")
		}
	}
}

//方法1性能好一点  （计算机中更适合做移位操作）
//[appuser@amyguo bitmap]$ go test -bench=. missingNum_test.go
//goos: linux
//goarch: amd64
//Benchmark_missingNumber  	100000000	        10.2 ns/op
//Benchmark_missingNumber2 	100000000	        10.1 ns/op
//PASS
//ok  	command-line-arguments	2.047s
//[appuser@amyguo bitmap]$ go test -bench=. missingNum_test.go
//goos: linux
//goarch: amd64
//Benchmark_missingNumber  	122079315	         9.93 ns/op
//Benchmark_missingNumber2 	119939608	         9.96 ns/op
//PASS
//ok  	command-line-arguments	4.412s
//[appuser@amyguo bitmap]$ go test -bench=. missingNum_test.go
//goos: linux
//goarch: amd64
//Benchmark_missingNumber  	122809938	         9.80 ns/op
//Benchmark_missingNumber2 	100000000	        10.1 ns/op
//PASS
//ok  	command-line-arguments	3.216s
