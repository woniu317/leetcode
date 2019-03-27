package main

import (
	"fmt"
	"math"
)

func removeDuplicates(nums []int) int {
	if nums == nil {
		return 0
	}
	faster, slower := 1, 0
	for faster < len(nums) {
		if nums[slower] != nums[faster] {
			slower++
			nums[slower] = nums[faster]
		}
		faster++
	}
	return slower + 1
}

func removeElement(nums []int, val int) int {
	slower, faster := 0, 0
	for faster < len(nums) {
		if nums[faster] != val {
			nums[slower] = nums[faster]
			slower++
		}
		faster++
	}
	return slower
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	index := -1
	isOk := true
	for i := 0; i < len(haystack); i++ {
		isOk = true
		if len(haystack)-i < len(needle) {
			return index
		}
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				isOk = false
				break
			}
		}
		if isOk {
			return i
		}
	}
	return index
}

func divide(dividend int, divisor int) int {
	resultSinal := 1
	result := 0
	if dividend == 0 {
		return 0
	}
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	if dividend < 0 {
		dividend = -dividend
		resultSinal = -resultSinal
	}
	if divisor < 0 {
		divisor = -divisor
		resultSinal = -resultSinal
	}
	sum := 0
	for sum <= dividend {
		sum += divisor
		result++
	}
	return resultSinal * (result - 1)
}

func main() {
	//	nums := []int{1, 1, 2}
	//	fmt.Println(removeDuplicates(nums))
	//	fmt.Println(nums)
	//	fmt.Println(removeElement(nums, 1))
	//	fmt.Println(nums)
	a := -2147483648 / (-1)
	fmt.Println(a)
	fmt.Println(divide(1, 1))
}
