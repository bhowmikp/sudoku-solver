/*
Package sudoku implements a library to solve soduku puzzles
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

const sudokuSpots = 81

// printTable takes in a string and prints a sudoku table
func printTable(text string) {
	// dimensions of the sudoku table
	const (
		rowNum   = 9
		blockNum = 3
	)

	for row := 0; row < rowNum; row++ {
		rowBlock := text[rowNum*row : rowNum*(row+1)]

		for block := 0; block < blockNum; block++ {
			if block < blockNum-1 {
				fmt.Print(rowBlock[blockNum*(block):blockNum*(block+1)] + "|")
			} else {
				fmt.Println(rowBlock[blockNum*(block) : blockNum*(block+1)])
			}
		}

		if (row > 0 && row < rowNum-1) && ((row+1)%3 == 0) {
			fmt.Println("-----------")
		}
	}
}

// convertCoorToArrLocation takes in string, and row, column info in array and
// gets value at its corresponding location inside the array
func convertCoorToArrLocation(text string, row, column int) (value string) {
	return value
}

// convertArrLocationToCoor takes in a string and its array location and
// gets value at its corresponding row and column inside array
func convertArrLocationToCoor(text string, location int) (value string) {
	return value
}

// findFirstEmpty finds the first location in the puzzle that is empty
// and returns the location inside the array
func findFirstEmpty(text string) (value string) {
	return value
}

// safeInColumn determines if a certain value is the only one of its
// kind inside the column
func safeInColumn(text, value string, column int) (status bool) {
	return status
}

// safeInRow determines if a certain value is the only one of its
// kind inside the row
func safeInRow(text, value string, row int) (status bool) {
	return status
}

// safeInBlock determines if a certain value is the only one of its
// kind in its 3x3 block
func safeInBlock(text, value string, row, column int) (status bool) {
	return status
}

// valueSafe determines if the value is safe to be placed in that
// location of the puzzle
func valueSafe(text, value string, row, column int) (status bool) {
	return status
}

func solvePuzzle(text string) string {
	return text
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

func main() {
	text := ""
	text = userInput(text)
}
