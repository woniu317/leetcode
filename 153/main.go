package main

func findMin(nums []int) int {
	return divMin(nums)
}

func divMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}
	l := len(nums)
	lmin := divMin(nums[:l/2])
	rmin := divMin(nums[l/2:])
	return minI(lmin, rmin)
}

func minI(i1, i2 int) int {
	if i1 < i2 {
		return i1
	}
	return i2
}
