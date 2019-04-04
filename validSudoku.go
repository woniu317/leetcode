package main

import "fmt"

func main() {
	//board := [][]byte{{5, 3, 0, 0, 7, 0, 0, 0, 0}, {6, 0, 0, 1, 9, 5, 0, 0, 0}, {0, 9, 8, 0, 0, 0, 0, 6, 0}, {8, 0, 0, 0, 6, 0, 0, 0, 3}, {4, 0, 0, 8, 0, 3, 0, 0, 1}, {7, 0, 0, 0, 2, 0, 0, 0, 6}, {0, 6, 0, 0, 0, 0, 2, 8, 0}, {0, 0, 0, 4, 1, 9, 0, 0, 5}, {0, 0, 0, 0, 8, 0, 0, 7, 9}}
	//fmt.Println(isValidSudoku(board))
	//board2 := [][]byte{{8, 3, 0, 0, 7, 0, 0, 0, 0}, {6, 0, 0, 1, 9, 5, 0, 0, 0}, {0, 9, 8, 0, 0, 0, 0, 6, 0}, {8, 0, 0, 0, 6, 0, 0, 0, 3}, {4, 0, 0, 8, 0, 3, 0, 0, 1}, {7, 0, 0, 0, 2, 0, 0, 0, 6}, {0, 6, 0, 0, 0, 0, 2, 8, 0}, {0, 0, 0, 4, 1, 9, 0, 0, 5}, {0, 0, 0, 0, 8, 0, 0, 7, 9}}
	//board2 := [][]byte{{'8', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	//fmt.Println(isValidSudoku(board2))
	//board3 := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	//fmt.Println(isValidSudoku(board3))

	//fmt.Println(isValidSudoku([][]byte{{'.', '.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '3', '.', '.'}, {'.', '.', '.', '1', '8', '.', '.', '.', '.'}, {'.', '.', '.', '7', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '1', '.', '9', '7', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '3', '6', '.', '1', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '2', '.'}}))
	fmt.Println(isValidSudoku([][]byte{{'.', '.', '.', '.', '5', '.', '.', '1', '.'}, {'.', '4', '.', '3', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '3', '.', '.', '1'}, {'8', '.', '.', '.', '.', '.', '.', '2', '.'}, {'.', '.', '2', '.', '7', '.', '.', '.', '.'}, {'.', '1', '5', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '2', '.', '.', '.'}, {'.', '2', '.', '9', '.', '.', '.', '.', '.'}, {'.', '.', '4', '.', '.', '.', '.', '.', '.'}}))
}

type Oper func(i, j *int)

func isValidSudoku(board [][]byte) bool {
	if !validRows(board) || !validCols(board) {
		return false
	}
	if !validBlocks(board) {
		return false
	}
	return true
}

func validBlocks(numbers [][]byte) bool {
	for i := 0; i < len(numbers); i += 3 {
		for j := 0; j < len(numbers); j += 3 {
			if !validBlock(numbers, i, j) {
				return false
			}
		}
	}
	return true
}

func validBlock(numbers [][]byte, startRow, startCol int) bool {
	i, j := startRow, startCol
	existData := make(map[byte]struct{}, 9)
	for ; i < startRow+3; i++ {
		for j = startCol; j < startCol+3; j++ {
			if data := numbers[i][j]; validData(existData, data) {
				existData[data] = struct{}{}
			} else {
				return false
			}
		}
	}
	return true
}

func validCols(numbers [][]byte) bool {
	oper := func(i, j *int) {
		*i++
	}
	for i := 0; i < len(numbers[0]); i++ {
		if !isValid(numbers, 0, i, oper) {
			return false
		}
	}
	return true
}

func validRows(numbers [][]byte) bool {
	oper := func(i, j *int) {
		*j++
	}
	for i := 0; i < len(numbers); i++ {
		if !isValid(numbers, i, 0, oper) {
			return false
		}
	}
	return true
}

func isValid(board [][]byte, rowStart, colStart int, oper Oper) bool {
	existData := make(map[byte]struct{}, 9)
	for ; validIndex(board, rowStart, colStart); oper(&rowStart, &colStart) {
		if data := board[rowStart][colStart]; validData(existData, data) {
			existData[data] = struct{}{}
		} else {
			return false
		}
	}
	return true
}

func validData(existData map[byte]struct{}, data byte) bool {
	if data == '.' {
		return true
	}
	_, ok := existData[data]
	return !ok
}

func validIndex(numbers [][]byte, row, col int) bool {
	return 0 <= row && row < len(numbers) && 0 <= col && col < len(numbers[row])
}
