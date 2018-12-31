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

// return if a character is a board valid character or not
func validBoardUnit(text rune) bool{
	validCharacters := map[rune]bool{'1': true, '2': true, '3': true, '4': true,
		                       '5': true, '6': true, '7': true, '8': true,
													 '9': true}

	return validCharacters[text]
}

// convertCoorToArrLocation takes in board as string, boards row and column and
// returns the value at the location
func getValueAtRowColumn(text string, row, column int) string {
	var location = columnNum * (row - 1) + (column - 1)

	return text[location:location+1]
}

// findFirstEmpty finds the first location in the puzzle that is empty
// and returns the location inside the array
func findFirstEmpty(text string) int {
	for i, letter := range text {
		if validBoardUnit(letter) == false {
			return i
		}
	}
	return -1
}

// safeInColumn determines if the value is safe to be placed in the column
func safeInColumn(board string, value rune, column int) bool {
	found := make(map[rune]bool)
	found[value] = true

	for row := 1; row <= rowNum; row++ {
		var valueAtLocation = []rune(getValueAtRowColumn(board, row, column))

		if validBoardUnit(valueAtLocation[0]){
			if found[valueAtLocation[0]] {
				return false
			}
			found[valueAtLocation[0]] = true
		}
	}

	return true
}

// safeInRow determines if the value is safe to be placed in the row
func safeInRow(board string, value rune, row int) bool {
	found := make(map[rune]bool)
	found[value] = true

	for column := 1; column <= columnNum; column++ {
		var valueAtLocation = []rune(getValueAtRowColumn(board, row, column))

		if validBoardUnit(valueAtLocation[0]) {
			if found[valueAtLocation[0]] {
				return false
			}
			found[valueAtLocation[0]] = true
		}
	}

	return true
}

// safeInBlock determines if there are duplicates of valid characters in
// 3x3 block
func safeInBlock(board string, value rune, row, column int) bool {
	found := make(map[rune]bool)
	found[value] = true

	row = (row - 1) - ((row - 1) % 3) + 1
	column = (column - 1) - ((column - 1) % 3) + 1

	for blockRow := 1; blockRow <= rowNum / 3; blockRow++ {
		for blockColumn := 1; blockColumn <= columnNum / 3; blockColumn++ {
			var valueAtLocation = []rune(getValueAtRowColumn(board, row, column))

			if validBoardUnit(valueAtLocation[0]) {
				if found[valueAtLocation[0]] {
					return false
				}
				found[valueAtLocation[0]] = true
			}
		}
	}

	return true
}

// valueSafe determines if the value is safe to be placed in that
// location of the puzzle
func valueSafe(text string, value rune, row, column int) (status bool) {
	return safeInColumn(text, value, column) && safeInRow(text, value, row) && safeInBlock(text, value, row, column)
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
