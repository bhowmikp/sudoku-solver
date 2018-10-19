package main

import (
	"bufio"
	"fmt"
	"os"
)

const sudokuSpots = 81

func printTable(text string) {
	const rowNum = 9
	const blockNum = 3

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
