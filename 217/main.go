package main

import "fmt"

func containsDuplicate(nums []int) bool {
	tmpMap := make(map[int]struct{})
	for _, k := range nums {
		if _, ok := tmpMap[k]; ok {
			return true
		}
		tmpMap[k] = struct{}{}
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
}
