package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const wordsFile string = "words.txt"

func ReadWordsFile() []string {
	byteData, err := ioutil.ReadFile(wordsFile)

	if err != nil {
		fmt.Println(err)
	}

	text := string(byteData)
	words := strings.Fields(text)

	return words
}

func main() {
	var words []string
	go func() {
		words = ReadWordsFile()
	}()
	dice := RollDice()

	fmt.Println("Welcome to the boggle solver/computer version!")
	fmt.Println("Please choose one of the following")
	fmt.Println("Solo mode (1), computer solve (2), input board (3)")
	fmt.Println()
	fmt.Print("Action > ")

	var action int = 2
	fmt.Scanln(&action)

	if action < 1 || action > 3 { // Action is not valid
		action = 2
		fmt.Println("\nYou did not enter a valid number. Defaulting to 2")
	}

	if action < 3 {
		fmt.Println("Below is the board")
		fmt.Println()
		PrintDice(dice)
		fmt.Println()
	}
}
