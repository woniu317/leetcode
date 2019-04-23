package main

import (
	"fmt"
)

func main() {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "*"))
	fmt.Println(isMatch("acdcb", "a*c?b"))
	fmt.Println(isMatch("", "?"))
}

const (
	Match    = 1
	NotMatch = 0
	Default  = -1
)

var result []int

func isMatch(s string, p string) bool {
	initArray(s, p)
	match(s, p)
	return getSlice(len(p)+1, len(s), len(p)) == Match
}

func match(s, p string) {
	pLenPlus := len(p) + 1
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(p); j++ {
			if p[j] == '?' || s[i] == p[j] {
				setSlice(pLenPlus, i+1, j+1, getSlice(pLenPlus, i, j))
			}
			if p[j] == '*' {
				if getSlice(pLenPlus, i, j+1) == Match || getSlice(pLenPlus, i, j) == Match || getSlice(pLenPlus, i+1, j) == Match {
					setSlice(pLenPlus, i+1, j+1, Match)
				}
			}
		}
	}
}

func initArray(s string, p string) {
	width := len(p) + 1
	result = *initSlice(len(s)+1, width)
	result[0] = Match
	for j := 0; j < len(p); j++ {
		if p[j] == '*' {
			setSlice(width, 0, j+1, getSlice(width, 0, j))
		}
	}
}

func initSlice(length, width int) *[]int {
	re := make([]int, length*width)
	return &re
}

func getSlice(width, row, col int) int {
	return result[row*width+col]
}

func setSlice(width, row, col, value int) {
	result[width*row+col] = value
}
