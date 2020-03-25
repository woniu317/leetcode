package main

import "fmt"

func main() {
	fmt.Println(minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
}

func minimumTotal(triangle [][]int) int {

	ret := make([]int, len(triangle[len(triangle)-1]))
	if len(triangle) == 0 {
		return 0
	}
	ret[0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		ret[i] = ret[i-1] + triangle[i][i]
		for j := i - 1; j >= 1; j-- {
			ret[j] = min(ret[j-1], ret[j]) + triangle[i][j]
		}
		ret[0] = ret[0] + triangle[i][0]
	}

	res := ret[0]
	for _, v := range ret {
		if res > v {
			res = v
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
