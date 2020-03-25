package main

import "fmt"

func main() {
	//grid := [][]byte{{1, 1, 1, 1, 0}, {1, 1, 0, 1, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 0, 0}}
	//grid := [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}
	//grid := [][]byte{{'1', '1', '1'}, {'0', '1', '0'}, {'1', '1', '1'}}
	grid := [][]byte{{'1'}, {'1'}}
	fmt.Println(numIslands(grid))
}

func numIslands(grid [][]byte) int {
	landNum := 0
	curLandNum := 1
	row := len(grid)
	if row == 0 {
		return landNum
	}
	col := len(grid[0])
	flag := make([]int, row*col)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if flag[i*col+j] == 0 && grid[i][j] == '1' {
				flag[i*col+j] = curLandNum
				landNum = curLandNum
				trans(row, col, i, j, curLandNum, &grid, &flag)
				curLandNum++
			}
		}
	}
	return landNum
}

func trans(rowNum, colNum, row, col, curlandNum int, grid *[][]byte, flag *[]int) {
	//左
	if col > 0 && (*grid)[row][col-1] == '1' && (*flag)[row*colNum+col-1] == 0 {
		(*flag)[row*colNum+col-1] = curlandNum
		trans(rowNum, colNum, row, col-1, curlandNum, grid, flag)
	}
	//右
	if col < colNum-1 && (*grid)[row][col+1] == '1' && (*flag)[row*colNum+col+1] == 0 {
		(*flag)[row*colNum+col+1] = curlandNum
		trans(rowNum, colNum, row, col+1, curlandNum, grid, flag)
	}
	//上
	if row > 0 && (*grid)[row-1][col] == '1' && (*flag)[(row-1)*colNum+col] == 0 {
		(*flag)[(row-1)*colNum+col] = curlandNum
		trans(rowNum, colNum, row-1, col, curlandNum, grid, flag)
	}
	//下
	if row < rowNum-1 && (*grid)[row+1][col] == '1' && (*flag)[(row+1)*colNum+col] == 0 {
		(*flag)[(row+1)*colNum+col] = curlandNum
		trans(rowNum, colNum, row+1, col, curlandNum, grid, flag)
	}

}
