package main

import "testing"

func TestUserInput1(t *testing.T) {
	input := "111111111222222222333333333444444444555555555666666666777777777888888888999999999"
	value := userInput(input)

	if value != input {
		t.Errorf("Value was incorrect, got: %s, want %s", value, input)
	}
}

func TestValidBoardLength1(t *testing.T) {
	input := "111111111222222222333333333444444444555555555666666666777777777888888888999999999"
	output := false

	if validBoardLength(input) == output {
		t.Errorf("Board does not have appropriate number of cells")
	}
}

func TestValidBoardLength2(t *testing.T) {
	input := "123456789"
	output := true

	if validBoardLength(input) == output {
		t.Errorf("Board does not have appropriate number of cells")
	}
}

func TestValidBoardUnit1(t *testing.T) {
	input := '1'
	output := false

	if validBoardUnit(input) == output {
		t.Errorf("%q is a valid character", input)
	}
}

func TestValidBoardUnit2(t *testing.T) {
	input := '@'
	output := true

	if validBoardUnit(input) == output {
		t.Errorf("%q is not a valid character", input)
	}
}

func TestValidBoardUnit3(t *testing.T) {
	input := 'a'
	output := true

	if validBoardUnit(input) == output {
		t.Errorf("%q is a not valid character", input)
	}
}

func TestGetValueAtRowColumn1(t *testing.T) {
	board := "111111111222222222333333333444444444555555555666666666777777777888888888999999999"
	row, column := 1, 1
	output := '1'

	if getValueAtRowColumn(board, row, column) != output {
		t.Errorf("%q is at the location", output)
	}
}

func TestGetValueAtRowColumn2(t *testing.T) {
	board := "111111111222222222333333333444444444555555555666666666777777777888888888999999999"
	row, column := 5, 5
	output := '5'

	if getValueAtRowColumn(board, row, column) != output {
		t.Errorf("%q is at the location", output)
	}
}

func TestGetValueAtRowColumn3(t *testing.T) {
	board := "111111111222222222333333333444444444555555555666666666777777777888888888999999999"
	row, column := 9, 9
	output := '9'

	if getValueAtRowColumn(board, row, column) != output {
		t.Errorf("%q is at the location", output)
	}
}

func TestFindFirstEmpty1(t *testing.T) {
	board := "111111111222222222333333333444444444555555555666666666777777777888888888999999999"
	output := -1

	if findFirstEmpty(board) != output {
		t.Errorf("There are no empty cells in the board")
	}
}

func TestFindFirstEmpty2(t *testing.T) {
	board := "a11111111222222222333333333444444444555555555666666666777777777888888888999999999"
	output := 0

	if findFirstEmpty(board) != output {
		t.Errorf("The empty cell is at %d", output)
	}
}

func TestFindFirstEmpty3(t *testing.T) {
	board := "11111111122222222233333333344444444455555555566666666677777777788888888899999999."
	output := 80

	if findFirstEmpty(board) != output {
		t.Errorf("The empty cell is at %d", output)
	}
}

func TestFindFirstEmpty4(t *testing.T) {
	board := "111111111222222222333333333z44444444555555555666.66666777777777888888888999999999"
	output := 27

	if findFirstEmpty(board) != output {
		t.Errorf("The empty cell is at %d", output)
	}
}

func TestConvertLinearLocationToRowColumn1(t *testing.T) {
	location := 0
	outputRow, outputColumn := 1, 1

	funcRow, funcColumn := convertLinearLocationToRowColumn(location)

	if funcRow != outputRow || funcColumn != outputColumn {
		t.Errorf("Location to row, column conversion not occuring correctly")
	}
}

func TestConvertLinearLocationToRowColumn2(t *testing.T) {
	location := 80
	outputRow, outputColumn := 9, 9

	funcRow, funcColumn := convertLinearLocationToRowColumn(location)

	if funcRow != outputRow || funcColumn != outputColumn {
		t.Errorf("Location to row, column conversion not occuring correctly")
	}
}

func TestConvertLinearLocationToRowColumn3(t *testing.T) {
	location := 21
	outputRow, outputColumn := 3, 4

	funcRow, funcColumn := convertLinearLocationToRowColumn(location)

	if funcRow != outputRow || funcColumn != outputColumn {
		t.Errorf("Location to row, column conversion not occuring correctly")
	}
}

func TestSetCellValue1(t *testing.T) {
	board := "111111111222222222333333333444444444555555555666666666777777777888888888999999999"
	value := '2'
	row, column := 1, 3
	outputBoard := "112111111222222222333333333444444444555555555666666666777777777888888888999999999"

	if setCellValue(board, value, row, column) != outputBoard {
		t.Errorf("Cell value is not set properly")
	}
}

func TestSetCellValue2(t *testing.T) {
	board := "111111111222222222333333333444444444555555555666666666777777777888888888999999999"
	value := '9'
	row, column := 5, 9
	outputBoard := "111111111222222222333333333444444444555555559666666666777777777888888888999999999"

	if setCellValue(board, value, row, column) != outputBoard {
		t.Errorf("Cell value is not set properly")
	}
}

func TestSetCellValue3(t *testing.T) {
	board := "111111111222222222333333333444444444555555555666666666777777777888888888999999999"
	value := '2'
	row, column := 8, 5
	outputBoard := "111111111222222222333333333444444444555555555666666666777777777888828888999999999"

	if setCellValue(board, value, row, column) != outputBoard {
		t.Errorf("Cell value is not set properly")
	}
}

func TestSetCellValue4(t *testing.T) {
	board := "111111111222222222333333333444444444555555555666666666777777777888888888999999999"
	value := '2'
	row, column := 9, 9
	outputBoard := "111111111222222222333333333444444444555555555666666666777777777888888888999999992"

	if setCellValue(board, value, row, column) != outputBoard {
		t.Errorf("Cell value is not set properly")
	}
}

func TestSafeInColumn1(t *testing.T) {
	board := ".11111111.22222222.33333333.44444444.55555555.66666666.77777777.88888888.99999999"
	value := '1'
	column := 1
	output := true

	if safeInColumn(board, value, column) != output {
		t.Errorf("Placing value:%q at column:%d in board:%s is a valid move", value, column, board)
	}
}

func TestSafeInColumn2(t *testing.T) {
	board := "1111.11112222.22223333.33334444.44445555.55556666.66667777.77778888.88889999.9999"
	value := '1'
	column := 5
	output := true

	if safeInColumn(board, value, column) != output {
		t.Errorf("Placing value:%q at column:%d in board:%s is a valid move", value, column, board)
	}
}

func TestSafeInColumn3(t *testing.T) {
	board := "1111.11112222222223333.33334444.44445555755556666.66667777677778888.8888999949999"
	value := '1'
	column := 5
	output := true

	if safeInColumn(board, value, column) != output {
		t.Errorf("Placing value:%q at column:%d in board:%s is a valid move", value, column, board)
	}

}

func TestSafeInColumn4(t *testing.T) {
	board := "1111.11112222222223333.33334444144445555755556666.66667777677778888.8888999949999"
	value := '1'
	column := 5
	output := false

	if safeInColumn(board, value, column) != output {
		t.Errorf("Placing value:%q at column:%d in board:%s is not a valid move", value, column, board)
	}
}

func TestSafeInRow1(t *testing.T) {
	board := ".........222222222333333333444444444555555555666666666777777777888888888999999999"
	value := '1'
	row := 1
	output := true

	if safeInRow(board, value, row) != output {
		t.Errorf("Placing value:%q at row:%d in board:%s is a valid move", value, row, board)
	}
}

func TestSafeInRow2(t *testing.T) {
	board := "111111111222222222333333333444444444.........666666666777777777888888888999999999"
	value := '1'
	row := 5
	output := true

	if safeInRow(board, value, row) != output {
		t.Errorf("Placing value:%q at row:%d in board:%s is a valid move", value, row, board)
	}
}

func TestSafeInRow3(t *testing.T) {
	board := "111111111222222222333333333444444444.2..7.6.4666666666777777777888888888999999999"
	value := '1'
	row := 5
	output := true

	if safeInRow(board, value, row) != output {
		t.Errorf("Placing value:%q at row:%d in board:%s is a valid move", value, row, board)
	}
}

func TestSafeInRow4(t *testing.T) {
	board := "111111111222222222333333333444444444.2.17.6.4666666666777777777888888888999999999"
	value := '1'
	row := 5
	output := false

	if safeInRow(board, value, row) != output {
		t.Errorf("Placing value:%q at row:%d in board:%s is not a valid move", value, row, board)
	}
}

func TestSafeInBlock1(t *testing.T) {
	board := "...111111...222222...333333444444444555555555666666666777777777888888888999999999"
	value := '1'
	row, column := 2, 2
	output := true

	if safeInBlock(board, value, row, column) != output {
		t.Errorf("Placing value:%q at row:%d, column:%d in board:%s is a valid move", value, row, column, board)
	}
}

func TestSafeInBlock2(t *testing.T) {
	board := "111111111222222222333333333444...444555...555666...666777777777888888888999999999"
	value := '1'
	row, column := 6, 4
	output := true

	if safeInBlock(board, value, row, column) != output {
		t.Errorf("Placing value:%q at row:%d, column:%d in board:%s is a valid move", value, row, column, board)
	}
}

func TestSafeInBlock3(t *testing.T) {
	board := "111111111222222222333333333444.2.444555.7.5556666.4666777777777888888888999999999"
	value := '1'
	row, column := 6, 5
	output := true

	if safeInBlock(board, value, row, column) != output {
		t.Errorf("Placing value:%q at row:%d, column:%d in board:%s is a valid move", value, row, column, board)
	}
}

func TestSafeInBlock4(t *testing.T) {
	board := "111111111222222222333333333444.2.44455517.5556666.4666777777777888888888999999999"
	value := '1'
	row, column := 5, 6
	output := false

	if safeInBlock(board, value, row, column) != output {
		t.Errorf("Placing value:%q at row:%d, column:%d in board:%s is not a valid move", value, row, column, board)
	}
}

func TestValueSafe1(t *testing.T) {
	board := ".6.14.57.9.......15.1...8..18...2.....46379.....8...42..6...2.48.......3.12.8..9."
	value := '2'
	row, column := 9, 9
	output := false

	if valueSafe(board, value, row, column) != output {
		t.Errorf("Placing value:%q at row:%d, column:%d in board:%s is not a valid move", value, row, column, board)
	}
}

func TestValueSafe2(t *testing.T) {
	board := ".6.14.57.9.......15.1...8..18...2.....46379.....8...42..6...2.48.......3.12.8..9."
	value := '1'
	row, column := 6, 5
	output := true

	if valueSafe(board, value, row, column) != output {
		t.Errorf("Placing value:%q at row:%d, column:%d in board:%s is a valid move", value, row, column, board)
	}
}

func TestSolvePuzzle1(t *testing.T) {
	board := ".6.14.57.9.......15.1...8..18...2.....46379.....8...42..6...2.48.......3.12.8..9."
	output := false
	outputStatus, _ := solvePuzzle(board)

	if outputStatus != output {
		t.Errorf("This board should not be solvable")
	}
}

func TestSolvePuzzle2(t *testing.T) {
	board := "..529.6......753.99...3...8.896.......79.28.......7.9.6...4...5..472......1.692.."
	output := true
	outputStatus, _ := solvePuzzle(board)

	if outputStatus != output {
		t.Errorf("This board should be solvable")
	}

}

func TestSolvePuzzle3(t *testing.T) {
	board := ".1.....9.4.......5.9316.....4..17.2.....5.....3.29..7.....7895.9.......8.2.....1."
	output := true
	outputStatus, _ := solvePuzzle(board)

	if outputStatus != output {
		t.Errorf("This board should be solvable")
	}

}

func TestSolvePuzzle4(t *testing.T) {
	board := "7..61...3.....9....2....19.97....8....24617....6....42.59....3....2.....1...38..5"
	output := true
	outputStatus, _ := solvePuzzle(board)

	if outputStatus != output {
		t.Errorf("This board should be solvable")
	}

}
