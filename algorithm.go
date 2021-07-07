package main

import "fmt"

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

// Checks to see if the word found is already in the wordFoundList
func Contains(array []string, test string) bool {
	for _, a := range array {
		if a == test {
			return true
		}
	}
	return false
}

// Recursively finds nearby board positions
// In boggle, you can move vertically, horizontally, and diagonally
func FindNearby(board [4][4]string, x int, y int, currentWord string) {
	searches++
	// NOTE: we use x and y here to easily determine location. Unlike
	// typical axis, y is in the first position because of the way the
	// cookie crumbles. Don't mess with it!!
	currentWord += board[y][x]

	// Here we define the rules of the game.
	// No words under three letters!
	if len(currentWord) >= 3 {
		wordResult := TestWord(currentWord)
		switch wordResult {
		case 0: // The word is valid
			if !Contains(wordsFoundList, currentWord) {
				fmt.Println(currentWord)
				wordsFoundList = append(wordsFoundList, currentWord)
				wordsFound++
				return
			}
		case 2: // Word and future words are invalid
			return
		}
	}

	// Can go west
	if x > 0 {
		FindNearby(board, x-1, y, currentWord)
	}
	// Can go east
	if x < 3 {
		FindNearby(board, x+1, y, currentWord)
	}
	// Can go north
	if y > 0 {
		FindNearby(board, x, y-1, currentWord)
	}
	// Can go south
	if y < 3 {
		FindNearby(board, x, y+1, currentWord)
	}

	// Can go southeast
	if x < 3 && y < 3 {
		FindNearby(board, x+1, y+1, currentWord)
	}
	// Can go southwest
	if x > 0 && y < 3 {
		FindNearby(board, x-1, y+1, currentWord)
	}
	// Can go northwest
	if x > 0 && y > 0 {
		FindNearby(board, x-1, y-1, currentWord)
	}
	// Can go northeast
	if x < 3 && y < 0 {
		FindNearby(board, x-1, y+1, currentWord)
	}
}
