package main

import "fmt"

func main() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[0][1] = "Y"
	board[0][2] = "Z"


	fmt.Println(board)
}