package main

import "fmt"

const (
	Match    = 1
	NotMatch = 0
	Default  = -1
)

var result [][]int

func isMatch(s string, p string) bool {
	initArray(s, p)
	return match(s, p, 0, 0)
}

func match(s, p string, sIndex, pIndex int) bool {
	matchResult := NotMatch
	fmt.Println(sIndex, pIndex)
	if result[sIndex][pIndex] == Default {
		if len(p) == pIndex {
			if sIndex == len(s) {
				matchResult = Match
			} else {
				matchResult = NotMatch
			}
		} else {
			firstMatch := false
			if sIndex < len(s) && (p[pIndex] == '.' || p[pIndex] == s[sIndex]) {
				firstMatch = true
			}
			if pIndex+1 < len(p) && p[pIndex+1] == '*' {
				if match(s, p, sIndex, pIndex+2) || (firstMatch && match(s, p, sIndex+1, pIndex)) {
					matchResult = Match
				}
			} else {
				if firstMatch && match(s, p, sIndex+1, pIndex+1) {
					matchResult = Match
				}
			}
		}
		result[sIndex][pIndex] = matchResult
	}

	return result[sIndex][pIndex] == Match
}

func initArray(s string, p string) {
	var re [][]int
	for i := 0; i <= len(s); i++ {
		var temp []int
		for j := 0; j <= len(p); j++ {
			temp = append(temp, -1)
		}
		re = append(re, temp)
	}
	result = re
}

func main() {
	fmt.Println(isMatch("asdb", ".*"))
}
