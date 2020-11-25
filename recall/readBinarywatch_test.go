package recall

import (
	"fmt"
	"testing"
)

/*
401. 二进制手表
二进制手表顶部有 4 个 LED 代表 小时（0-11），底部的 6 个 LED 代表 分钟（0-59）。
每个 LED 代表一个 0 或 1，最低位在右侧。


H;8 4 2 1
      . .
M:32 16 8 4 2 1
      . .     .

例如，上面的二进制手表读取 “3:25”。
给定一个非负整数 n 代表当前 LED 亮着的数量，返回所有可能的时间。


示例：

输入: n = 1
返回: ["1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"]

提示：
输出的顺序没有要求。
小时不会以零开头，比如 “01:00” 是不允许的，应为 “1:00”。
分钟必须由两位数组成，可能会以零开头，比如 “10:2” 是无效的，应为 “10:02”。
超过表示范围（小时 0-11，分钟 0-59）的数据将会被舍弃，也就是说不会出现 "13:00", "0:61" 等时间。
*/

func readBinaryWatch(num int) []string {
	res := []string{}
	// 假想 10个 bit 灯 index  --->       h1=>0   h2=>1   h4=>2   h8=>3   m1=>4   m2=>5   m4=>6   m8=>7   m16=>8   m32=>9
	//start = 0 回溯从bit 的 index = 0 开始, 10 代表时分灯的总数, 相当于10个灯中选中num个亮的灯
	backTrack(0, 10, num, []int{}, &res)
	return res
}

//backTrack 回溯算法 cap 灯的总数量, target 目标亮灯的数量 path []int  亮灯的index集合(记录回溯的路径) res 指针返回结果
func backTrack(start, cap, target int, path []int, res *[]string) {
	if len(path) == target { //选择亮灯的数量达到target
		min, hour := 0, 0
		for _, v := range path {
			if v >= 4 { // 如果灯的下表超过3 则表示这些灯是 minute的灯
				min += 1 << uint((v - 4))
			} else { //如果 灯的index 小于4 表示是 hour的灯亮
				hour += 1 << uint(v)
			}
		}
		if min < 60 && hour < 12 { //排除不符合规则的,判断min hour 值是否有效
			*res = append(*res, fmt.Sprintf("%d:%02d", hour, min)) //格式化拼接成字符串
		}
		path = []int{} //重置 回溯的path
		return
	}
	for i := start; i < cap; i++ {
		path = append(path, i)
		backTrack(i+1, cap, target, path, res) //注意这里是 i+1
		path = path[:len(path)-1]              //对 path 进行rollback
	}
}

//func readBinaryWatch2(num int) []string {
//	res := make([]string, 0)
//	total := 10
//	var backTrack func(start int, path []int)
//	backTrack = func(start int, path []int) {
//		if len(path) == num {
//			min, hour := 0, 0
//			for _, v := range path {
//				if v >= 4 {
//					min += 1 << uint(v-4)
//				} else {
//					hour += 1 << uint(v)
//				}
//			}
//			if min < 60 && hour < 12 {
//				res = append(res, fmt.Sprintf("%d:%02d", hour, min))
//			}
//			path = []int{}
//			return
//		}
//		for i := start; i < total; i++ {
//			path = append(path, i)
//			backTrack(i+1, path)
//			path = path[:len(path)-1]
//		}
//	}
//
//	backTrack(0, []int{})
//	return res
//}

//暴力法：
func readBinaryWatch3(num int) []string {
	ret := make([]string, 0)
	for i := 0; i < 12; i++ {
		for j := 0; j <= 59; j++ {
			if bitNums(i)+bitNums(j) == num {
				ret = append(ret, fmt.Sprintf("%d:%02d", i, j))
			}
		}
	}
	return ret
}

func bitNums(i int) int {
	num := 0
	for i > 0 {
		if i%2 == 1 {
			num++
		}
		i = i >> 1
	}
	return num
}

func Test_readBinaryWatch(t *testing.T) {
	fmt.Println(readBinaryWatch(1))
}
