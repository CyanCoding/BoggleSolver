package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func RollDice() (dice [5][5]string) {
	allDiceValues := FillDice()

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			rand.Seed(time.Now().UnixNano())
			// Gets a random number between 0 and 6
			randomNumber := rand.Intn(6)
			value := allDiceValues[i][randomNumber]

			dice[i][j] = value
		}
	}

	return
}

func FillDice() (d [25][6]string) {
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

	d[16][0] = "e"
	d[16][1] = "x"
	d[16][2] = "l"
	d[16][3] = "d"
	d[16][4] = "r"
	d[16][5] = "i"

	d[17][0] = "h"
	d[17][1] = "o"
	d[17][2] = "p"
	d[17][3] = "c"
	d[17][4] = "s"
	d[17][5] = "a"

	d[18][0] = "e"
	d[18][1] = "x"
	d[18][2] = "l"
	d[18][3] = "d"
	d[18][4] = "r"
	d[18][5] = "i"

	d[19][0] = "h"
	d[19][1] = "o"
	d[19][2] = "p"
	d[19][3] = "c"
	d[19][4] = "s"
	d[19][5] = "a"

	d[20][0] = "e"
	d[20][1] = "x"
	d[20][2] = "l"
	d[20][3] = "d"
	d[20][4] = "r"
	d[20][5] = "i"

	d[21][0] = "h"
	d[21][1] = "o"
	d[21][2] = "p"
	d[21][3] = "c"
	d[21][4] = "s"
	d[21][5] = "a"

	d[22][0] = "e"
	d[22][1] = "x"
	d[22][2] = "l"
	d[22][3] = "d"
	d[22][4] = "r"
	d[22][5] = "i"

	d[23][0] = "h"
	d[23][1] = "o"
	d[23][2] = "p"
	d[23][3] = "c"
	d[23][4] = "s"
	d[23][5] = "a"

	d[24][0] = "e"
	d[24][1] = "x"
	d[24][2] = "l"
	d[24][3] = "d"
	d[24][4] = "r"
	d[24][5] = "i"

	return
}

func ManuallySetDice() (dice [5][5]string) {
	fmt.Println("Fill in each of the five rows below with spaces between letters")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("(1) > ")
	row1, _ := reader.ReadString('\n')

	fmt.Print("(2) > ")
	row2, _ := reader.ReadString('\n')

	fmt.Print("(3) > ")
	row3, _ := reader.ReadString('\n')

	fmt.Print("(4) > ")
	row4, _ := reader.ReadString('\n')

	fmt.Print("(5) > ")
	row5, _ := reader.ReadString('\n')

	fmt.Println()

	row1 = strings.TrimSuffix(row1, "\n")
	row2 = strings.TrimSuffix(row2, "\n")
	row3 = strings.TrimSuffix(row3, "\n")
	row4 = strings.TrimSuffix(row4, "\n")
	row5 = strings.TrimSuffix(row5, "\n")

	row1Array := strings.Fields(row1)
	row2Array := strings.Fields(row2)
	row3Array := strings.Fields(row3)
	row4Array := strings.Fields(row4)
	row5Array := strings.Fields(row5)

	dice[0][0] = row1Array[0]
	dice[0][1] = row1Array[1]
	dice[0][2] = row1Array[2]
	dice[0][3] = row1Array[3]
	dice[0][4] = row1Array[4]

	dice[1][0] = row2Array[0]
	dice[1][1] = row2Array[1]
	dice[1][2] = row2Array[2]
	dice[1][3] = row2Array[3]
	dice[1][4] = row2Array[4]

	dice[2][0] = row3Array[0]
	dice[2][1] = row3Array[1]
	dice[2][2] = row3Array[2]
	dice[2][3] = row3Array[3]
	dice[2][4] = row3Array[4]

	dice[3][0] = row4Array[0]
	dice[3][1] = row4Array[1]
	dice[3][2] = row4Array[2]
	dice[3][3] = row4Array[3]
	dice[3][4] = row4Array[4]

	dice[4][0] = row5Array[0]
	dice[4][1] = row5Array[1]
	dice[4][2] = row5Array[2]
	dice[4][3] = row5Array[3]
	dice[4][4] = row5Array[4]

	return dice
}

func PrintDice(dice [5][5]string) {
	if width == 4 {
		fmt.Println(dice[0][0], "\t", dice[0][1], "\t", dice[0][2], "\t", dice[0][3])
		fmt.Println(dice[1][0], "\t", dice[1][1], "\t", dice[1][2], "\t", dice[1][3])
		fmt.Println(dice[2][0], "\t", dice[2][1], "\t", dice[2][2], "\t", dice[2][3])
		fmt.Println(dice[3][0], "\t", dice[3][1], "\t", dice[3][2], "\t", dice[3][3])
	} else if width == 5 {
		fmt.Println(dice[0][0], "\t", dice[0][1], "\t", dice[0][2], "\t", dice[0][3], "\t", dice[0][4])
		fmt.Println(dice[1][0], "\t", dice[1][1], "\t", dice[1][2], "\t", dice[1][3], "\t", dice[1][4])
		fmt.Println(dice[2][0], "\t", dice[2][1], "\t", dice[2][2], "\t", dice[2][3], "\t", dice[2][4])
		fmt.Println(dice[3][0], "\t", dice[3][1], "\t", dice[3][2], "\t", dice[3][3], "\t", dice[3][4])
		fmt.Println(dice[4][0], "\t", dice[4][1], "\t", dice[4][2], "\t", dice[4][3], "\t", dice[4][4])
	}
}
