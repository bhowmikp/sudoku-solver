/*
Package sudoku implements a library to solve soduku puzzles
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	blockNum = 3
	rowNum = 9
	columnNum = 9
	sudokuSpots = 81
)

// printTable takes in a string and prints a sudoku table
// to print to stdout pass os.Stdout
func printTable(w io.Writer, text string) {
	for row := 0; row < rowNum; row++ {
		rowBlock := text[rowNum*row : rowNum*(row+1)]

		for block := 0; block < blockNum; block++ {
			if block < blockNum-1 {
				fmt.Fprint(w, rowBlock[blockNum*(block):blockNum*(block+1)] + "|")
			} else {
				fmt.Fprintln(w, rowBlock[blockNum*(block) : blockNum*(block+1)])
			}
		}

		if (row > 0 && row < rowNum-1) && ((row+1)%3 == 0) {
			fmt.Fprintln(w, "-----------")
		}
	}
}

// convertCoorToArrLocation takes in board as string, boards row and column and
// returns the value at the location
func getValueAtRowColumn(text string, row, column int) (value string) {
	return value
}

// findFirstEmpty finds the first location in the puzzle that is empty
// and returns the location inside the array
func findFirstEmpty(text string) int {
	numbers := map[rune]bool{'0': true, '1': true, '2': true, '3': true,
													 '4': true, '5': true, '6': true, '7': true,
													 '8': true, '9': true}

	for i, letter := range text {
		if numbers[letter] == false {
			return i
		}
	}
	return -1
}

// safeInColumn determines if a certain value is the only one of its
// kind inside the column
func safeInColumn(board, value string, column int) (status bool) {
	return status
}

// safeInRow determines if a certain value is the only one of its
// kind inside the row
func safeInRow(board, value string, row int) (status bool) {
	return status
}

// safeInBlock determines if a certain value is the only one of its
// kind in its 3x3 block
func safeInBlock(board, value string, row, column int) (status bool) {
	return status
}

// valueSafe determines if the value is safe to be placed in that
// location of the puzzle
func valueSafe(text, value string, row, column int) (status bool) {
	return status
}

// solves the sudoku puzzle
func solvePuzzle(text string) string {
	return text
}

// check if a board is valid length
func validBoardLength(text string) bool {
	return len(text) == sudokuSpots
}

// userInput takes a text and checks if the text is valid
// if the text is not keeps asking for input till user gives valid text
func userInput(text string) string {
	reader := bufio.NewReader(os.Stdin)

	for len(text) != sudokuSpots {
		fmt.Println("Enter the sudoku table in one line (81 characters):")
		text, _ = reader.ReadString('\n')
	}

	return text
}
