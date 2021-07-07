package main

// Recursively finds nearby board positions
// In boggle, you can move vertically, horizontally, and diagonally
func FindNearby(board [4][4]string, x int, y int, currentWord string) {
	// NOTE: we use x and y here to easily determine location. Unlike
	// typical axis, y is in the first position because of the way the
	// cookie crumbles. Don't mess with it!!
	currentWord += board[y][x]

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
