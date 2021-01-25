package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	ret := make([]int, len(triangle))
	for i, v := range triangle[len(triangle)-1] {
		ret[i] = v
	}

	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			ret[j] = min(ret[j], ret[j+1]) + triangle[i][j]
		}
	}

	return ret[0]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
}
