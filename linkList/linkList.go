package main

import "fmt"

func main()  {
	nums := []int{1,4,6,2,9,3,5}
	fmt.Println("bubble:",BubbleSort(nums))
	fmt.Println("quick:",QuickSort(nums,0, len(nums)-1))
}

func BubbleSort(nums []int)[]int  {
	for i:= 0 ;i< len(nums);i++{
		for j:= i+1;j<len(nums);j++{
			if nums[i] > nums[j] {
				nums[i],nums[j] = nums[j],nums[i]
			}
		}
	}
	return nums
}

func QuickSort(nums []int,left,right int) []int {
	if left < right {
		flag := nums[left]
		i := left
		j := right
		for i < j {
			if nums[i] < flag {
				i++
			}
			if nums[j] > flag {
				j--
			}
			nums[i],nums[j] = nums[j],nums[i]
		}
		QuickSort(nums,left,i-1)
		QuickSort(nums,j+1,right)
	}
	return nums
}

func HeapSort()  {
	
}
