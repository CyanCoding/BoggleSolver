package main

import (
	"fmt"
	"math/rand"
	"time"
)

func RollDice() (dice [4][4]string) {
	allDiceValues := FillDice()

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			rand.Seed(time.Now().UnixNano())
			// Gets a random number between 0 and 6
			randomNumber := rand.Intn(6)
			value := allDiceValues[i][randomNumber]

			dice[i][j] = value
		}
	}

	return
}

func FillDice() (d [16][6]string) {
	d[0][0] = "z"
	d[0][1] = "l"
	d[0][2] = "n"
	d[0][3] = "r"
	d[0][4] = "h"
	d[0][5] = "n"

	d[1][0] = "qu"
	d[1][1] = "n"
	d[1][2] = "m"
	d[1][3] = "i"
	d[1][4] = "h"
	d[1][5] = "v"

	d[2][0] = "d"
	d[2][1] = "s"
	d[2][2] = "y"
	d[2][3] = "i"
	d[2][4] = "t"
	d[2][5] = "t"

	d[3][0] = "w"
	d[3][1] = "a"
	d[3][2] = "o"
	d[3][3] = "t"
	d[3][4] = "o"
	d[3][5] = "t"

	d[4][0] = "m"
	d[4][1] = "i"
	d[4][2] = "o"
	d[4][3] = "t"
	d[4][4] = "c"
	d[4][5] = "u"

	d[5][0] = "f"
	d[5][1] = "a"
	d[5][2] = "k"
	d[5][3] = "p"
	d[5][4] = "s"
	d[5][5] = "f"

	d[6][0] = "t"
	d[6][1] = "y"
	d[6][2] = "r"
	d[6][3] = "l"
	d[6][4] = "t"
	d[6][5] = "e"

	d[7][0] = "v"
	d[7][1] = "e"
	d[7][2] = "e"
	d[7][3] = "s"
	d[7][4] = "i"
	d[7][5] = "n"

	d[8][0] = "o"
	d[8][1] = "b"
	d[8][2] = "j"
	d[8][3] = "a"
	d[8][4] = "b"
	d[8][5] = "o"

	d[9][0] = "i"
	d[9][1] = "e"
	d[9][2] = "s"
	d[9][3] = "t"
	d[9][4] = "o"
	d[9][5] = "s"

	d[10][0] = "n"
	d[10][1] = "a"
	d[10][2] = "e"
	d[10][3] = "a"
	d[10][4] = "g"
	d[10][5] = "e"

	d[11][0] = "e"
	d[11][1] = "h"
	d[11][2] = "n"
	d[11][3] = "e"
	d[11][4] = "g"
	d[11][5] = "w"

	d[12][0] = "h"
	d[12][1] = "r"
	d[12][2] = "v"
	d[12][3] = "t"
	d[12][4] = "w"
	d[12][5] = "e"

	d[13][0] = "d"
	d[13][1] = "l"
	d[13][2] = "y"
	d[13][3] = "e"
	d[13][4] = "r"
	d[13][5] = "u"

	d[14][0] = "e"
	d[14][1] = "x"
	d[14][2] = "l"
	d[14][3] = "d"
	d[14][4] = "r"
	d[14][5] = "i"

	d[15][0] = "h"
	d[15][1] = "o"
	d[15][2] = "p"
	d[15][3] = "c"
	d[15][4] = "s"
	d[15][5] = "a"

	return
}

func PrintDice(dice [4][4]string) {
	fmt.Println(dice[0][0], dice[0][1], dice[0][2], dice[0][3])
	fmt.Println(dice[1][0], dice[1][1], dice[1][2], dice[1][3])
	fmt.Println(dice[2][0], dice[2][1], dice[2][2], dice[2][3])
	fmt.Println(dice[3][0], dice[3][1], dice[3][2], dice[3][3])
}
