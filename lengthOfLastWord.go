package main

import "fmt"

func main() {
	fmt.Println(lengthOfLastWord("b a  "))
}

func lengthOfLastWord(s string) int {
	result := 0
	temp := 0
	for _, ss := range s {
		if ('a' <= ss && 'z' >= ss) || ('A' <= ss && 'Z' >= ss) {
			temp++
		} else {
			if temp > 0 {
				result = temp
			}
			temp = 0
		}
	}
	if temp > 0 {
		return temp
	}
	return result
}
