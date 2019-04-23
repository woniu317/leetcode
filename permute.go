package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println(permute([]int{1}))
}

var result [][]int

func permute(nums []int) [][]int {
	result = nil
	permuteByStart(nums, 0)
	return result
}

func permuteByStart(nums []int, start int) {
	if start == len(nums)-1 {
		re := make([]int, len(nums))
		copy(re, nums)
		result = append(result, re)
		return
	}
	for i := start; i < len(nums); i++ {
		nums[i], nums[start] = nums[start], nums[i]
		permuteByStart(nums, start+1)
		nums[i], nums[start] = nums[start], nums[i]
	}
}

func permuteUnique(nums []int) [][]int {
	result = nil
	permuteUniqueByStart(nums, 0)
	return result
}

func permuteUniqueByStart(nums []int, start int) {
	if start == len(nums)-1 {
		re := make([]int, len(nums))
		copy(re, nums)
		result = append(result, re)
		return
	}
	for i := start; i < len(nums); i++ {
		if i == start || valid(nums, start, i) {
			nums[i], nums[start] = nums[start], nums[i]
			permuteUniqueByStart(nums, start+1)
			nums[i], nums[start] = nums[start], nums[i]
		}
	}
}

func valid(nums []int, i, j int) bool {
	target := nums[j]
	for ; i < j; i++ {
		if nums[i] == target {
			return false
		}
	}
	return true
}
