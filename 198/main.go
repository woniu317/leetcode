package main

import "fmt"

func main() {
	nums := []int{2, 7, 9, 3, 1}
	fmt.Println(rob(nums))
}

func rob(nums []int) int {
	lenth := len(nums)
	if lenth < 1 {
		return 0
	} else if lenth == 1 {
		return nums[0]
	}
	preMax := nums[0]
	curMax := maxI(nums[0], nums[1])
	for i := 2; i < lenth; i++ {
		tmpMax := curMax
		curMax = maxI(curMax, preMax+nums[i])
		preMax = tmpMax
	}
	return curMax
}

func maxI(a, b int) int {
	if a > b {
		return a
	}
	return b
}
