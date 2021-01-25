package main

import (
	"fmt"
	"math"
)

func main() {
	r := maxProductPath([][]int{{-1, -2, 3, -3, -3, 3, 1, 0, 4, 1, 4, -3, 2, 3, -2},
		{-4, 0, 4, 4, 1, 1, 0, -2, -4, 1, 4, 1, -1, -4, 2},
		{-3, -2, 2, -4, -1, -1, -2, 0, 0, -1, -4, 2, 4, -2, 0},
		{2, 2, 3, 3, -2, 4, -3, -3, 3, -4, 0, -1, 4, 4, 4},
		{-4, 0, -3, -2, -3, -2, 2, -1, 2, 4, 3, 4, -4, -3, -1},
		{-4, 0, -3, 0, 2, 2, 0, 3, 3, -1, 3, 3, 1, -3, -3},
		{-1, -3, 4, 3, 4, -2, -3, -3, -4, -3, 0, 2, -4, 4, -4},
		{3, -1, 3, -4, -1, 1, 0, -3, 2, 1, -4, -2, 3, 0, 3},
		{3, -1, -2, 2, -4, 4, -1, 3, -3, 2, -4, 0, -2, -2, 1},
		{4, 1, -3, 2, -1, 2, 4, 0, 3, 4, 3, 2, 2, 3, 2},
		{-2, 1, 0, 3, -3, -4, -4, 4, 1, -2, -1, 3, 3, -3, 3},
		{0, -2, 0, 2, 2, -3, -4, 2, 1, 2, -2, 0, -1, -3, 4},
		{4, -1, -2, 3, 2, -1, -4, -4, 2, -3, 1, 0, 0, 4, -3},
		{-4, 2, -1, 4, 4, -2, -4, 4, -1, 1, -4, -3, -4, 3, 0},
		{2, -1, -4, -2, 0, -3, 3, -4, 2, 0, 2, -3, 0, -3, 4}})
	fmt.Println(r)
	fmt.Println(CheckPhoneNoLength("111111111111"))
}

func IsLongerThan(str string, length int) bool {
	return len(str) > length
}

func CheckPhoneNoLength(PhoneNo string) bool {
	return !IsLongerThan(PhoneNo, 15)
}

func maxProductPath(grid [][]int) int {
	var (
		nr  int = len(grid)
		nc  int = len(grid[0])
		ans int = math.MinInt64
		dfs func(int, int, int)
	)
	dfs = func(row, col, value int) {
		value *= grid[row][col]
		if value == 0 || row == nr-1 && col == nc-1 {
			if value > ans {
				ans = value
			}
			return
		}
		if col < nc-1 {
			dfs(row, col+1, value)
		}
		if row < nr-1 {
			dfs(row+1, col, value)
		}
	}
	dfs(0, 0, 1)
	if ans < 0 {
		return -1
	} else {
		return ans % 1000000007
	}
}
