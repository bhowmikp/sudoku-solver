package main

import "testing"

func TestUserInput(t *testing.T) {
	input := "111111111222222222333333333444444444555555555666666666777777777888888888999999999"
	value := userInput(input)

	if value != input {
		t.Errorf("Value was incorrect, got: %s, want %s", value, input)
	}
}
