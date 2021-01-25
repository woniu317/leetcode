package main

import (
	"fmt"
	"time"
)

func main() {
	// longestCommonSubsequence("abcde", "ace")
	fmt.Println(time.ParseDuration("10min"))
}

func longestCommonSubsequence(text1 string, text2 string) int {
	if text1 == "" || text2 == "" {
		return 0
	}

	length1 := len(text1)
	length2 := len(text2)
	data := make([][]int, length2+1)

	for i := 0; i < length2+1; i++ {
		data[i] = make([]int, length1+1)
	}

	for i := 0; i < length2; i++ {
		for j := 0; j < length1; j++ {
			if text2[i] == text1[j] {
				data[i+1][j+1] = data[i][j] + 1
				continue
			}
			data[i+1][j+1] = max(data[i][j+1], data[i+1][j])
		}
	}

	return data[length2][length1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
