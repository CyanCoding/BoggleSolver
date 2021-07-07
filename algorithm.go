package main

import "fmt"

func TestWord(currentWord string) int {
	for i := 0; i < len(words); i++ {
		if len(words[i]) != len(currentWord) { // Word is not proper length
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
			// We've gotten past the length of the word without
			// hitting an incorrect letter. This means there COULD
			// be potential solutions, but we haven't found one yet
			if len(words[i]) > len(currentWord) {
				return 1 // Potential for future match
			}

			if words[i][j] != currentWord[j] { // Match failed, try next word
				break
			}
		}
	}

	// We've gotten this far and haven't returned, so return bad
	return 2 // FAIL, no match or future matches with available letters
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
			fmt.Println(currentWord)
			wordsFound++
			return
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
