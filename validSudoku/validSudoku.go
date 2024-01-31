package main

import (
	"fmt"
)


func isValidSudoku(board [][]byte) bool { 

/*
    var rows, cols, square [9][9]bool

    for i, row := range board {
        for j, v :=  range row {
            if v != '.' {
                k := int(v)-49
                if rows[i][k] || cols[j][k] || square[i/3*3 + j/3][k] {
                    return false
                }

                rows[i][k], cols[j][k], square[i/3*3 + j/3][k] = true, true, true
            }
        }
    }
    return true
*/


var rows, cols, square []map[int]bool

for i, row := range board {
    for j, v := range row {
        for v != '.' {
            k := int(v)
            if rows[i][k] || cols[j][k] || square[i/3*3 + j/3][k] {
                return false
            }
            rows[i][k], cols[j][k], square[i/3*3 + j/3][k] = true, true, true
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
        {'.', '.', '.', '.', '8', '.', '.', '.', '9'},
    }

	fmt.Println(isValidSudoku(sudokuBoard))

}
