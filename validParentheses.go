package main

import "fmt"

var matchParenthesesMap = map[string]string{
	"(": ")",
	"{": "}",
	"[": "]",
}

func isStartParentheses(s string) bool {
	switch s {
	case "(":
		return true
	case "[":
		return true
	case "{":
		return true
	default:
		return false
	}
}

func isValid(s string) bool {
	lenth := len(s)
	unValid := make([]int, 0, lenth)
	for i := 0; i < lenth; i++ {
		if isStartParentheses(s[i : i+1]) {
			unValid = append(unValid, i)
			continue
		}
		if len(unValid) == 0 {
			return false
		}
		unvalideIndex := unValid[len(unValid)-1]
		if matchParenthesesMap[s[unvalideIndex:unvalideIndex+1]] != s[i:i+1] {
			return false
		}
		unValid = unValid[:len(unValid)-1]
	}
	return len(unValid) == 0
}
func main() {
	fmt.Println(isValid("(())"))
	fmt.Println(isValid("()[]{}"))
}
