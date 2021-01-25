package main

import "fmt"

func main() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	fmt.Println(maximalSquareDp(matrix))
}

func isOne(a byte) bool {
	return a == '1'
}
func maximalSquareDp(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	rowLen := len(matrix)
	colLen := len(matrix[0])
	dp := make([][]int, rowLen)
	curLen := 0
	for i := range dp {
		dp[i] = make([]int, colLen)
	}
	for i := 0; i < colLen; i++ {
		if isOne(matrix[0][i]) {
			dp[0][i] = 1
			curLen = 1
		}
	}
	for i := 0; i < rowLen; i++ {
		if isOne(matrix[i][0]) {
			dp[i][0] = 1
			curLen = 1
		}
	}

	for i := 1; i < rowLen; i++ {
		for j := 1; j < colLen; j++ {
			if !isOne(matrix[i][j]) {
				continue
			}
			dp[i][j] = min(dp[i-1][j-1], dp[i][j-1], dp[i-1][j]) + 1
			curLen = max(curLen, dp[i][j])
		}
	}
	return curLen * curLen
}

func min(a, b, c int) int {
	if a > b {
		if b > c {
			return c
		}
		return b
	}
	if a > c {
		return c
	}
	return a
}

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	rowLen := len(matrix)
	colLen := len(matrix[0])
	curArea := 0
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			if matrix[i][j] == '1' {
				curArea = max(curArea, subMaximalSquare(matrix, i, j))
			}
		}
	}
	return curArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func subMaximalSquare(matrix [][]byte, row, col int) int {
	rowLen := len(matrix)
	colLen := len(matrix[0])
	curLen := 0
	curRow := row
	curCol := col
are:
	for curRow < rowLen && curCol < colLen {
		for i := row; i <= curRow; i++ {
			if matrix[i][curCol] != '1' {
				break are
			}
		}
		for i := col; i <= curCol; i++ {
			if matrix[curRow][i] != '1' {
				break are
			}
		}
		curLen++
		curRow++
		curCol++
	}
	return curLen * curLen
}
