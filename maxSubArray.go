package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSubArray([]int{1, 2, 3, 4, 5}))
}

func maxSubArray(nums []int) int {
	curSum := 0
	res := math.MinInt64
	for _, num := range nums {
		curSum = max(curSum+num, num)
		res = max(res, curSum)
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
