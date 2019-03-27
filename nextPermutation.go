package main

import "fmt"

func nextPermutation(nums []int) {
	index := findReplacePosition(nums)
	if index == -1 {
		reverse(nums)
		return
	}
	rIndex := findMinGreaterByValueIndex(nums, index)
	swap(nums, index, rIndex)
	reverseFromIndex(nums, index+1)
}

func reverseFromIndex(nums []int, index int) {
	i := index
	j := len(nums) - 1
	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
}

func swap(nums []int, index1 int, index2 int) {
	nums[index1], nums[index2] = nums[index2], nums[index1]
}

//从index位置起查找nums中比nums[index]大的元素下标
func findMinGreaterByValueIndex(nums []int, index int) int {
	value := nums[index]
	result := index + 1
	for i := index + 1; i < len(nums); i++ {
		if nums[i] > value {
			result = i
		} else {
			break
		}
	}
	return result
}

//逆置数组
func reverse(nums []int) {
	i := 0
	j := len(nums) - 1
	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
}

//从后往前找到最后一个非增序的元素坐标
func findReplacePosition(nums []int) int {
	i := len(nums) - 1
	j := i - 1
	for j >= 0 {
		if nums[j] >= nums[i] {
			j--
			i--
		} else {
			break
		}
	}
	return j
}

func main() {
	nums := []int{5, 1, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}
