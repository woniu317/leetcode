package main

import "fmt"

func main() {
	fmt.Println(countBinarySubstrings("100111001"))
	fmt.Println(countBinarySubstrings("00110"))
}

func countBinarySubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}
	lastNum := 0
	nowNum := 0
	curch := uint8(' ')
	ret := 0
	for i := range s {
		if s[i] == curch {
			nowNum++
		} else {
			ret += min(lastNum, nowNum)
			curch = s[i]
			lastNum = nowNum
			nowNum = 1
		}
	}
	ret += min(lastNum, nowNum)
	return ret
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
