package main

import (
	"fmt"
)


func isValidSudoku(board [][]byte) bool { 

    var rows, columns, squares [9][9]bool
    for i, row := range board {
        for j, v := range row {
            if v != '.' {
                k := int(v)-49
                if rows[i][k] || columns[j][k] || squares[i/3*3 + j/3][k] {
                    return false
                }
                rows[i][k], columns[j][k], squares[i/3*3 + j/3][k] = true, true, true
            }
        }
    }
    return true
}

func main() {

	sudokuBoard := [][]byte{
        {'5', '3', '.', '.', '7', '.', '.', '.', '.'},
        {'6', '.', '.', '1', '9', '5', '.', '.', '.'},
        {'.', '9', '8', '.', '.', '.', '.', '6', '.'},
        {'8', '.', '.', '.', '6', '.', '.', '.', '3'},
        {'4', '.', '.', '8', '.', '3', '.', '.', '1'},
        {'7', '.', '.', '.', '2', '.', '.', '.', '6'},
        {'.', '6', '.', '.', '.', '.', '2', '8', '.'},
        {'.', '.', '.', '4', '1', '9', '.', '.', '5'},
        {'7', '.', '.', '.', '8', '.', '.', '7', '9'},
    }

	fmt.Println(isValidSudoku(sudokuBoard))

}
