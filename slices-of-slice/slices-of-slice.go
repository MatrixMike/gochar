package main

import (
	"fmt"
	"strings"
)
// possibly confusing strings and runes here 
func main() {
	// Create a tic-tac-toe board.
	board := [][]string{				// board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"qwqw"},
	}
/*
	board2 := [][]string {
	board2[1][1] = "Z"
	}
*/
	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
		board[2][3] = "Z"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], "")) // both blank and empty string work here 
	}
		fmt.Printf("Type: %T Value: %v\n", board,board)
}
