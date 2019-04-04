package main

import (
	"fmt"
)

func main() {
	//board := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	board := [][]byte{{'.', '.', '9', '7', '4', '8', '.', '.', '.'}, {'7', '.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '2', '.', '1', '.', '9', '.', '.', '.'}, {'.', '.', '7', '.', '.', '.', '2', '4', '.'}, {'.', '6', '4', '.', '1', '.', '5', '9', '.'}, {'.', '9', '8', '.', '.', '.', '3', '.', '.'}, {'.', '.', '.', '8', '.', '3', '.', '2', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.', '6'}, {'.', '.', '.', '2', '7', '5', '9', '.', '.'}}
	printB(board)
	solveSudoku(board)
	printB(board)

}

func printB(board [][]byte) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			fmt.Print(string(board[i][j]) + " ")
		}
		fmt.Println()
	}
	fmt.Println("***********************************")
}

type Mark struct {
	row [9]int
	col [9]int
	box [9]int
}

type Oper func(board [][]byte, row, col *int) bool

func solveSudoku(board [][]byte) {
	oper := func(board [][]byte, row, col *int) bool {
		if *row >= len(board) {
			return false
		}
		if *col < len(board[*row])-1 {
			*col++
			return true
		}
		if *row < len(board)-1 {
			*row++
			*col = 0
			return true
		}
		return false
	}
	mark := new(Mark)
	initMark(board, mark)
	solveSudokuOne(board, 0, 0, oper, mark)
}

func initMark(board [][]byte, mark *Mark) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '.' {
				continue
			}
			n := uint(board[i][j] - '0')
			markInfo(i, j, n, mark)
		}
	}
}

//找到结果时返回结果
func solveSudokuOne(board [][]byte, row, col int, oper Oper, mark *Mark) bool {
	if board[row][col] != '.' {
		if oper(board, &row, &col) {
			return solveSudokuOne(board, row, col, oper, mark)
		}
		return true
	}
	nextValue := uint(0)
	for findNextValue(row, col, &nextValue, mark) {
		r, c := row, col
		board[r][c] = byte(nextValue) + '0'
		markInfo(row, col, nextValue, mark)
		if oper(board, &r, &c) {
			solved := solveSudokuOne(board, r, c, oper, mark)
			if solved {
				return true
			}
			remarkInfo(row, col, nextValue, mark)
		} else {
			return true
		}
	}
	board[row][col] = '.'
	return false
}

func remarkInfo(row, col int, value uint, mark *Mark) {
	markInfo(row, col, value, mark)
}

func markInfo(row, col int, value uint, mark *Mark) {
	n := 1 << (value - 1)
	mark.row[row] ^= n
	mark.col[col] ^= n
	mark.box[row/3*3+col/3] ^= n
}

//查找比nextValue大的下一个可以用的value，若存在则返回true，value存储next值
func findNextValue(row, col int, value *uint, mark *Mark) bool {
	for i := *value + 1; i < 10; i++ {
		n := 1 << (i - 1)
		if mark.row[row]^n >= mark.row[row] &&
			mark.col[col]^n >= mark.col[col] &&
			mark.box[row/3*3+col/3]^n >= mark.box[row/3*3+col/3] {
			*value = i
			return true
		}
	}
	return false
}
