package main

import "fmt"

func main() {
	fmt.Println(longestValidParentheses("(()"))
	fmt.Println(longestValidParentheses(")()())"))
}

//312ms
func longestValidParentheses(s string) int {
	maxLen := 0
	stack := make([]int, 0, len(s)+1)
	stack = append(stack, -1)
	for j := 0; j < len(s); j++ {
		if s[j:j+1] == "(" {
			stack = append(stack, j)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, j)
			} else {
				if matchLen := j - stack[len(stack)-1]; matchLen > maxLen {
					maxLen = matchLen
				}
			}
		}
	}
	return maxLen
}

func longestValidParentheses1(s string) int {
	maxLen := 0
	for i := 0; i < len(s); i++ {
		if s[i:i+1] == ")" {
			continue
		}
		tempMatch := 0
		remaindLeft := 0
		for j := i; j < len(s); j++ {
			if s[j:j+1] == ")" {
				remaindLeft--
			} else {
				remaindLeft++
			}
			if remaindLeft == 0 {
				tempMatch = j - i + 1
			} else if remaindLeft < 0 {
				i = j - 1
				break
			}
			if tempMatch > maxLen {
				maxLen = tempMatch
			}

		}
	}
	return maxLen
}
