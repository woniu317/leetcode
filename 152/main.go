package main

import "fmt"

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
}
func maxProduct(nums []int) int {
	max := nums[0]
	imin := nums[0]
	imax := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			imin, imax = imax, imin
		}
		imax = maxI(imax*nums[i], nums[i])
		imin = minI(imin*nums[i], nums[i])
		max = maxI(imax, max)
	}
	return max
}
func maxI(i1, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}

func minI(i1, i2 int) int {
	if i1 < i2 {
		return i1
	}
	return i2
}
