package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}
func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	cols := len(grid[0])
	row := len(grid)
	dp := make([]int, cols)
	for i := 0; i < cols; i++ {
		dp[i] = math.MaxInt64
	}
	dp[0] = 0
	for i := 0; i < row; i++ {
		dp[0] = dp[0] + grid[i][0]
		for j := 1; j < cols; j++ {
			dp[j] = min(dp[j-1], dp[j]) + grid[i][j]
		}
	}
	return dp[cols-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
