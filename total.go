package main

import (
	"fmt"
)

//长度为n的数组中所有数都在0-n-1范围内，找出重复的数
func FindRepeatedNum(nums []int)int  {
	for i:= 0;i<len(nums);i++{
		for (i !=nums[i]) {
			if( nums[i] == nums[nums[i]]){
				return nums[i]
			}
			nums[i],nums[nums[i]] = nums[nums[i]],nums[i]
		}
	}
	return 0
}

//给定一个包含 n + 1 个整数的数组 nums ，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数。
func GetDuplicate(nums []int) int {
	slow,fast := nums[0],nums[nums[0]]
	for (slow != fast) {
		slow,fast = nums[slow],nums[nums[fast]]
	}
	fast = 0
	for (slow != fast){
		slow,fast = nums[slow],nums[fast]
	}
	return slow
}

//在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
func findNumberIn2DArray(matrix [][]int, target int) bool {
	//以左下角为原点
	i:=len(matrix)-1//获取右下角y坐标
	j:=0//获取右下角x坐标
	for i>-1{
		if j<len(matrix[i]){
			if target<matrix[i][j]{
				i--  //小于target,向上查找
			}else if target>matrix[i][j]{
				j++ //大于targat,向右查找
			}else if target==matrix[i][j]{
				return true
			}
		}else{
			return false//超出数组返回false
		}
	}
	return false//超出matrix返回false
}

func quickSort(arr []int,left,right int)  {
	if left < right {
		flag := arr[left]
		i := left
		j := right
		fmt.Println(flag,i,j)
		for i < j {
			for arr[i] < flag {
				i++
			}

			for arr[j] > flag {
				j--
			}
			arr[i], arr[j] = arr[j], arr[i]
		}
		quickSort(arr, left, i-1)
		quickSort(arr, j+1, right)
	}
}

func main()  {
	//fmt.Println(FindRepeatedNum([]int{2,3,1,0,2,5,3}))
	//fmt.Println(GetDuplicate([]int{2,3,5,4,3,2,7,6}))
	nums := []int{41, 24, 76, 11, 45, 64, 21, 69, 19, 19, 39}
	quickSort(nums,0,len(nums)-1)
	fmt.Println(nums)
}

