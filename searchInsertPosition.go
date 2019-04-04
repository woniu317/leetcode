package main

import "fmt"

func searchInsert(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start < end {
		mid := (start + end) / 2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	if nums[start] < target {
		return start + 1
	}
	return start
}

func main() {
	//fmt.Println(searchInsert([]int{0, 2, 3}, 1))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
}
