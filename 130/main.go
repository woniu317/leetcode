package main

func main() {
	// board := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
	board := [][]byte{{'O', 'O'}, {'O', 'O'}}
	solve(board)
}

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	flag := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		flag[i] = make([]bool, len(board[i]))
	}
	row := len(board)
	col := len(board[0])
	for i := 0; i < len(board[0]); i++ {
		if board[0][i] == 'O' {
			markFlag(board, 0, i, flag)
		}
		if board[row-1][i] == 'O' {
			markFlag(board, row-1, i, flag)
		}
		if board[i][0] == 'O' {
			markFlag(board, i, 0, flag)
		}
		if board[i][col-1] == 'O' {
			markFlag(board, i, col-1, flag)
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if !flag[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}

func markFlag(board [][]byte, i, j int, flag [][]bool) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return
	}
	if board[i][j] == 'X' || flag[i][j] {
		return
	}
	flag[i][j] = true
	markFlag(board, i-1, j, flag)
	markFlag(board, i+1, j, flag)
	markFlag(board, i, j-1, flag)
	markFlag(board, i, j+1, flag)
}
