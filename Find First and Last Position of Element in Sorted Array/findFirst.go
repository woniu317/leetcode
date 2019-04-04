package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{1}, 1))
}

//查找target在有序数组nums里面的区间
func searchRange(nums []int, target int) []int {
	first := findFirstPosition(nums, target)
	if first == -1 {
		return []int{-1, -1}
	}
	last := findLastPosition(nums, target)
	return []int{first, last}
}

//查找与target相同的最后一个元素
func findLastPosition(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	start, end := 0, len(nums)-1
	mid := 0
	for start < end {
		mid = (start + end) / 2
		if nums[mid] <= target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	if start < len(nums) && nums[start] == target {
		return start
	}
	if start > 0 && nums[start-1] == target {
		return start - 1
	}
	return -1
}

//查找与target相同的第一个元素
func findFirstPosition(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	start, end := 0, len(nums)-1
	mid := 0
	for start < end {
		mid = (start + end) / 2
		if nums[mid] >= target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	if end >= 0 && nums[end] == target {
		return end
	}
	if end < len(nums)-1 {
		if nums[end+1] == target {
			return end + 1
		}
	}
	return -1
}
