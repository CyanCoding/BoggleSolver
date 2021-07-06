package main

import "fmt"

func main() {
	dice := RollDice()

	fmt.Println("Welcome to the boggle solver/computer version!")
	fmt.Println("Please choose one of the following")
	fmt.Println("Solo mode (1), computer solve (2), input board (3)")
	fmt.Print("Action > ")

	var action int = 2
	fmt.Scanln(&action)

	if action < 1 || action > 3 { // Action is not valid
		action = 2
		fmt.Println("You did not enter a valid number. Defaulting to 2")
	}

	fmt.Println("Below is the board")
	PrintDice(dice)
}
