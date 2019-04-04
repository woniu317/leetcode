package main

import "fmt"

func search(nums []int, target int) int {

	if len(nums) == 0 {
		return -1
	}
	var start, end = 0, len(nums) - 1
	var mid int
	for end-start > 1 {
		mid = (start + end) / 2
		if targetInMidBefore(nums, start, end, mid, target) {
			end = mid
		} else if targetInMidAfter(nums, start, end, mid, target) {
			start = mid
		} else {
			return -1
		}
	}
	if nums[start] == target {
		return start
	} else if nums[end] == target {
		return end
	}
	return -1
}

func targetInMidBefore(nums []int, start, end, mid, target int) bool {
	if nums[start] < nums[mid] && target <= nums[mid] && target >= nums[start] {
		return true
	}
	if nums[start] > nums[mid] && (target < nums[mid] || target >= nums[start]) {
		return true
	}
	return false
}

func targetInMidAfter(nums []int, start, end, mid, target int) bool {
	if nums[mid] > nums[end] && (nums[end] >= target || target >= nums[mid]) {
		return true
	}
	if nums[mid] < nums[end] && nums[end] >= target && nums[mid] <= target {
		return true
	}
	return false
}

func main() {
	fmt.Println(search([]int{2, 3, 0, 1}, 0))
	fmt.Println(search([]int{2, 3, 0, 1}, 1))
	fmt.Println(search([]int{2, 3, 0, 1}, 2))
	fmt.Println(search([]int{2, 3, 0, 1}, 3))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}
