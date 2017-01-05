package main

import (
	"fmt"
	"strings"
)
// possibly confusing strings and runes here 
func main() {
		var a [3]string
	a[1] = "one"
	// Create a tic-tac-toe board.
	board := [][]string{				// board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"qwqw"},
		// want to initialise next row to some part of a from above
	}

/*
	board2 := [][]string {
	board2[1][1] = "Z"
	}
*/
	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[2] = []string{"qq"}
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
// 		board[2][3] = "Z"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], "")) // both blank and empty string work here 
	}
		fmt.Printf("Type: %T Value: %v\n", board,board)
				fmt.Printf("Type: %T Value: %v\n", a,a)
}
