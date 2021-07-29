package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var diceOptions = [150]string{
	"z",
	"l",
	"n",
	"r",
	"h",
	"n",
	"qu",
	"n",
	"m",
	"i",
	"h",
	"v",
	"d",
	"s",
	"y",
	"i",
	"t",
	"t",
	"w",
	"a",
	"o",
	"t",
	"o",
	"t",
	"m",
	"i",
	"o",
	"t",
	"c",
	"u",
	"f",
	"a",
	"k",
	"p",
	"s",
	"f",
	"t",
	"y",
	"r",
	"l",
	"t",
	"e",
	"v",
	"e",
	"e",
	"s",
	"i",
	"n",
	"o",
	"b",
	"j",
	"a",
	"b",
	"o",
	"i",
	"e",
	"s",
	"t",
	"o",
	"s",
	"n",
	"a",
	"e",
	"a",
	"g",
	"e",
	"e",
	"h",
	"n",
	"e",
	"g",
	"w",
	"h",
	"r",
	"v",
	"t",
	"w",
	"e",
	"d",
	"l",
	"y",
	"e",
	"r",
	"u",
	"e",
	"x",
	"l",
	"d",
	"r",
	"i",
	"h",
	"o",
	"p",
	"c",
	"s",
	"a",
	"e",
	"x",
	"l",
	"d",
	"r",
	"i",
	"h",
	"o",
	"p",
	"c",
	"s",
	"a",
	"e",
	"x",
	"l",
	"d",
	"r",
	"i",
	"h",
	"o",
	"p",
	"c",
	"s",
	"a",
	"e",
	"x",
	"l",
	"d",
	"r",
	"i",
	"h",
	"o",
	"p",
	"c",
	"s",
	"a",
	"e",
	"x",
	"l",
	"d",
	"r",
	"i",
	"h",
	"o",
	"p",
	"c",
	"s",
	"a",
	"e",
	"x",
	"l",
	"d",
	"r",
	"i",
}

var DiceMap = map[string]int{}

func RollDice() (dice [5][5]string) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			rand.Seed(time.Now().UnixNano())
			// Gets a random number between 0 and 150
			randomNumber := rand.Intn(150)
			value := diceOptions[randomNumber]
			time.Sleep(1 * time.Millisecond)

			dice[i][j] = value
			DiceMap[value] = i + j
		}
	}

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
