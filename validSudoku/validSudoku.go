package main

import (
	"fmt"
)

func isValidSudoku(board [][]byte) bool {

	col := [9][9]bool{}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {

			if col[i][board[i][j]] {
				return false
			}
			col[i][board[i][j]] = true
		}
	}

	return true
}

func main() {

	sudokuMap := [][]byte{
		{1, 2, 3, 8, 4, 6, 7, 8, 9},
	}

	fmt.Println(isValidSudoku(sudokuMap))

}
