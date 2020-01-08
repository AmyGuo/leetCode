package _81_findUnsortedSubarray

//给定一个整数数组，你需要寻找一个连续的子数组，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。
//
//你找到的子数组应是最短的，请输出它的长度。
//
//示例 1:
//
//输入: [2, 6, 4, 8, 10, 9, 15]
//输出: 5
//解释: 你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
//说明 :
//
//输入的数组长度范围在 [1, 10,000]。
//输入的数组可能包含重复元素 ，所以升序的意思是<=

//先排序，然后做对比，记录位置相同的元素，要找的连续子数组就是长度-相同元素的数量（比较low）
func FindUnsortedSubarray(nums []int) int {
	if len(nums) == 0 || len(nums) > 10000 {
		return 0
	}

	unSortNums := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		unSortNums = append(unSortNums, nums[i])
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	count := 0
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == unSortNums[i] {
			count++
			index++
		}
		if index-i != 1 {
			break
		}
	}

	newIndex := len(nums) - 1
	for i := len(nums) - 1; i > index; i-- {
		if nums[i] == unSortNums[i] {
			count++
			newIndex--
		}
		if i-newIndex != 1 {
			break
		}
	}
	return len(nums) - count
}

//利用选择排序的思想，找到左右位置不正确的index临界值，两个的差值就是要找的长度
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func FindUnsortedSubarray2(nums []int) int {
	l := len(nums)
	r := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				r = max(r, j)
				l = min(l, i)
			}
		}
	}

	if r-1 < 0 {
		return 0
	} else {
		return r - l + 1
	}
}

//这个方法背后的想法仍然是选择排序。我们需要找到无序子数组中最小元素和最大元素分别对应的正确位置，来求得我们想要的无序子数组的边界。
//
//为了达到这一目的，此方法中，我们使用 栈 。我们从头遍历 numsnums 数组，如果遇到的数字大小一直是升序的，我们就不断把对应的下标压入栈中，这么做的目的是因为这些元素在目前都是处于正确的位置上。一旦我们遇到前面的数比后面的数大，也就是 nums[j]nums[j] 比栈顶元素小，我们可以知道 nums[j]nums[j] 一定不在正确的位置上。
//
//为了找到 nums[j]nums[j] 的正确位置，我们不断将栈顶元素弹出，直到栈顶元素比 nums[j]nums[j] 小，我们假设栈顶元素对应的下标为 kk ，那么我们知道 nums[j]nums[j] 的正确位置下标应该是 k + 1k+1 。
//
//我们重复这一过程并遍历完整个数组，这样我们可以找到最小的 kk， 它也是无序子数组的左边界。
//
//类似的，我们逆序遍历一遍 numsnums 数组来找到无序子数组的右边界。这一次我们将降序的元素压入栈中，如果遇到一个升序的元素，我们像上面所述的方法一样不断将栈顶元素弹出，直到找到一个更大的元素，以此找到无序子数组的右边界。

//速度最快的方法：利用栈和选择排序的思想
type stack []interface{}

func (s *stack) Push(value interface{}) {
	*s = append(*s, value)
}

func (s *stack) Pop() interface{} {
	if s.Len() == 0 {
		return nil
	}
	temp := *s
	value := temp[temp.Len()-1]
	*s = temp[:temp.Len()-1]
	return value
}
func (s *stack) Top() interface{} {
	temp := *s
	return temp[temp.Len()-1]
}

func (s *stack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *stack) Clear() {
	*s = nil
}
func (s *stack) Len() int {
	return len(*s)
}
func FindUnsortedSubarray3(nums []int) int {
	stk := new(stack)
	l := len(nums) - 1
	r := 0
	stk.Push(0)
	for i := 1; i < len(nums); i++ {
		for !stk.IsEmpty() && nums[(stk.Top()).(int)] > nums[i] {
			l = min(l, (stk.Pop()).(int))
		}
		stk.Push(i)
	}
	stk.Clear()
	stk.Push(len(nums) - 1)
	for i := len(nums) - 2; i >= 0; i-- {
		for !stk.IsEmpty() && nums[(stk.Top()).(int)] < nums[i] {
			r = max(r, (stk.Pop()).(int))

		}
		stk.Push(i)
	}
	if r-l > 0 {
		return r - l + 1
	} else {
		return 0
	}

}

//别人写的，和方法3思想一致
func FindUnsortedSubarray4(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	stack := make([]int, n)
	count := 0
	min := n - 1
	stack[count] = 0
	count++
	j := 0
	for i := 1; i < n; i++ {
		for count > 0 {
			j = stack[count-1]
			if nums[j] > nums[i] {
				if min > j {
					min = j
				}
				count--
			} else {
				break
			}
		}
		stack[count] = i
		count++
	}

	stack[0] = n - 1
	count = 1
	j = 0
	max := 0
	for i := n - 2; i >= 0; i-- {
		for count > 0 {
			j = stack[count-1]
			if nums[j] < nums[i] {
				if max < j {
					max = j
				}
				count--
			} else {
				break
			}
		}
		stack[count] = i
		count++
	}

	if min > max {
		return 0
	}
	return (max - min + 1)
}
