package main

import "fmt"

func main() {
	//fmt.Println(uniquePaths(7, 3))
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
}

func uniquePaths(m int, n int) int {
	x, y := 1, 1
	minV := min(m, n)
	for i, j := m+n-2, 1; j < minV; i-- {
		x *= i
		j += 1
	}
	for i := 2; i < minV; i++ {
		y *= i
	}
	fmt.Println(x, y)
	return x / y
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || obstacleGrid[0][0] == 1 {
		return 0
	}
	dp := make([]int, len(obstacleGrid[0]))
	dp[0] = 1
	row := len(obstacleGrid)
	cols := len(obstacleGrid[0])
	for i := 0; i < row; i++ {
		for j := 0; j < cols; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
			} else if j > 0 {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[cols-1]
}
