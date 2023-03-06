package main

import (
	"errors"
	"fmt"
) // ValidRule checks if the current cell is valid
func ValidRule(board [][]byte, j, i byte) bool {
	num := board[j][i]
	// Check if the number is valid in the row
	for p := byte(0); p < 9; p++ {
		if board[j][p] == num && p != i {
			return false
		}
	}
	// Check if the number is valid in the column
	for p := byte(0); p < 9; p++ {
		if board[p][i] == num && p != j {
			return false
		}
	}
	// Check if the number is valid in the 3x3 grid(printing grid)
	for p := byte(0); p < 3; p++ {
		for l := byte(0); l < 3; l++ {
			if board[j-j%3+p][i-i%3+l] == num && (j-j%3+p != j || i-i%3+l != i) {
				return false
			}
		}
	}
	return true
}

// findNextEmpty returns the next empty cell.
func findNextEmpty(board [][]byte) (byte, byte, error) {
	for j := byte(0); j < 9; j++ {
		for i := byte(0); i < 9; i++ {
			if board[j][i] == '.' {
				return j, i, nil
			}
		}
	}
	return byte(0), byte(0), errors.New("no empty cell to fill in")
}
func solveSudoku(board [][]byte) {
	j, i, e := findNextEmpty(board)
	if e != nil {
		return
	}
	var backtracing func(byte, byte) bool
	backtracing = func(j, i byte) bool {
		// fill in avabilable choices
		for p := byte(1); p <= 9; p++ {
			board[j][i] = 48 + p
			if ValidRule(board, j, i) {
				j, i, e := findNextEmpty(board)
				if e != nil {
					return true
				}
				status := backtracing(j, i)
				if status {
					return true
				}
			}
			board[j][i] = '.'
		}
		return false
	}
	// start backtracking
	backtracing(j, i)
}

// print board
func printBoard(board [][]byte) {
	for _, row := range board {
		for ix, col := range row {
			if ix == 8 {
				fmt.Printf("%c\n", col)
			} else {
				fmt.Printf("%c ", col)
			}
		}
	}
}
func main() {
	board := [][]byte{
		{'.', '9', '6', '.', '4', '.', '.', '.', '1'},
		{'1', '.', '.', '.', '6', '.', '.', '.', '4'},
		{'5', '.', '4', '8', '1', '.', '3', '9', '.'},
		{'.', '.', '7', '9', '5', '.', '.', '4', '3'},
		{'.', '3', '.', '.', '8', '.', '.', '.', '.'},
		{'4', '.', '5', '.', '2', '3', '.', '1', '8'},
		{'.', '1', '.', '6', '3', '.', '.', '5', '9'},
		{'.', '5', '9', '.', '7', '.', '8', '3', '.'},
		{'.', '.', '3', '5', '9', '.', '.', '.', '7'},
	}
	solveSudoku(board)
	printBoard(board)
}
