package main

import "strconv"

func TestWord(currentWord string) int {
	futureSuccess := false
	for i := 0; i < len(words); i++ {
		wordsChecked++
		if len(words[i]) < len(currentWord) { // Word is not proper length
			continue
		}

		if words[i] == currentWord {
			return 0 // YES! Match
		}

		// Tests each letter to see if there are future possibilities

		// BE CAREFUL, we use len(words[i]), which might be
		// longer than len(currentWord) in some instances
		// Make sure you break before you hit an error
		for j := 0; j < len(words[i]); j++ {
			if j+1 == len(currentWord) && words[i][j] == currentWord[j] { // Hit the word's limit
				// Might have a succesful match in the future
				futureSuccess = true
				break
			}
			if words[i][j] != currentWord[j] { // Match failed, try next word
				break
			}
		}
	}
	if futureSuccess {
		return 1
	} else {
		return 2
	}
}

// Checks to see if a string is in a string array
func containsString(array []string, test string) bool {
	for _, a := range array {
		if a == test {
			return true
		}
	}
	return false
}

// Check to see if an int is in an int array
// Used for determining if a dice code is in diceCodesUsed
func ContainsInt(array []int, test int) bool {
	for _, a := range array {
		if a == test {
			return true
		}
	}
	return false
}

// Recursively finds nearby board positions
// In boggle, you can move vertically, horizontally, and diagonally
func FindNearby(board [5][5]diceValue, x int, y int, currentWord string, repeatLetters bool, diceCodesUsed []int) {
	searches++
	// NOTE: we use x and y here to easily determine location. It might
	// seem strange, but we use y as the first value since it determines
	// the "height" (row) of the array, which in real life would be
	// the y-value.
	value := board[y][x]

	currentWord += value.character

	// We make sure that the dice code isn't already present in this word
	if !repeatLetters {
		if ContainsInt(diceCodesUsed, value.id) {
			diceCodesUsed = nil
			return
		} else { // It does contain the letter
			diceCodesUsed = append(diceCodesUsed, value.id)
		}
	}

	// Here we define the rules of the game.
	// No words under three letters!
	if len(currentWord) >= 3 {
		wordResult := TestWord(currentWord)
		switch wordResult {
		case 0: // The word is valid
			if !containsString(wordsFoundList, currentWord) {
				s.SetMessage(ColorCyan + "Finding matches (" + strconv.Itoa(percentDone) +
					"% • " + strconv.Itoa(wordsFound) + ")...")
				wordsFoundList = append(wordsFoundList, currentWord)
				wordsFound++
			}
		case 2: // Word and future words are invalid
			diceCodesUsed = nil
			return
		}
	}

	// Can go west
	if x > 0 {
		FindNearby(board, x-1, y, currentWord, repeatLetters, diceCodesUsed)
	}
	// Can go east
	if x < (width - 1) {
		FindNearby(board, x+1, y, currentWord, repeatLetters, diceCodesUsed)
	}
	// Can go north
	if y > 0 {
		FindNearby(board, x, y-1, currentWord, repeatLetters, diceCodesUsed)
	}
	// Can go south
	if y < (width - 1) {
		FindNearby(board, x, y+1, currentWord, repeatLetters, diceCodesUsed)
	}

	// Can go southeast
	if x < (width-1) && y < (width-1) {
		FindNearby(board, x+1, y+1, currentWord, repeatLetters, diceCodesUsed)
	}
	// Can go southwest
	if x > 0 && y < (width-1) {
		FindNearby(board, x-1, y+1, currentWord, repeatLetters, diceCodesUsed)
	}
	// Can go northwest
	if x > 0 && y > 0 {
		FindNearby(board, x-1, y-1, currentWord, repeatLetters, diceCodesUsed)
	}
	// Can go northeast
	if x < (width-1) && y > 0 {
		FindNearby(board, x+1, y-1, currentWord, repeatLetters, diceCodesUsed)
	}
}
