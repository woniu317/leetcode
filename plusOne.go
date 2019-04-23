package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{9}))
}

func plusOne(digits []int) []int {
	c := 1
	for i := len(digits) - 1; i >= 0; i-- {
		s := digits[i] + c
		digits[i] = s % 10
		c = s / 10
	}
	if c == 1 {
		return append([]int{1}, digits...)
	}
	return digits
}
