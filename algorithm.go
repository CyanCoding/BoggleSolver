package main

// Recursively finds nearby board positions
func FindNearby(board [4][4]string, x int, y int) {

	// Can go west
	if x > 0 {
		FindNearby(board, x-1, y)
	}
	// Can go east
	if x < 3 {
		FindNearby(board, x+1, y)
	}
	// Can go north
	if y > 0 {
		FindNearby(board, x, y-1)
	}
	// Can go south
	if y < 3 {
		FindNearby(board, x, y+1)
	}

	// Can go southeast
	if x < 3 && y < 3 {
		FindNearby(board, x+1, y+1)
	}
	// Can go southwest
	if x > 0 && y < 3 {
		FindNearby(board, x-1, y+1)
	}
	// Can go northwest
	if x > 0 && y > 0 {
		FindNearby(board, x-1, y-1)
	}
	// Can go northeast
	if x < 3 && y < 0 {
		FindNearby(board, x-1, y+1)
	}
}
