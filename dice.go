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

// Used for determining whether we have a duplicate ID or not
var diceIDsUsed = []int{}

type diceValue struct {
	character string
	id        int
}

func RollDice() (dice [5][5]diceValue) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			rand.Seed(time.Now().UnixNano())
			// Gets a random number between 0 and 150
			randomNumber := rand.Intn(150)
			value := diceOptions[randomNumber]
			time.Sleep(time.Millisecond)

			for {
				if ContainsInt(diceIDsUsed, randomNumber) {
					time.Sleep(time.Millisecond)
					randomNumber = rand.Intn(150)
				} else {
					val := diceValue{value, randomNumber}

					dice[i][j] = val
					break
				}
			}
		}
	}

	return
}

func ManuallySetDice() (dice [5][5]diceValue) {
	fmt.Println("Fill in each of the rows below with spaces between letters")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("(1) > ")
	row1, _ := reader.ReadString('\n')

	fmt.Print("(2) > ")
	row2, _ := reader.ReadString('\n')

	fmt.Print("(3) > ")
	row3, _ := reader.ReadString('\n')

	fmt.Print("(4) > ")
	row4, _ := reader.ReadString('\n')

	var row5 string
	var row5Array []string

	if width == 5 {
		fmt.Print("(5) > ")
		row5, _ = reader.ReadString('\n')
	}

	fmt.Println()

	row1 = strings.TrimSuffix(row1, "\n")
	row2 = strings.TrimSuffix(row2, "\n")
	row3 = strings.TrimSuffix(row3, "\n")
	row4 = strings.TrimSuffix(row4, "\n")
	if width == 5 {
		row5 = strings.TrimSuffix(row5, "\n")
	}

	row1Array := strings.Fields(row1)
	row2Array := strings.Fields(row2)
	row3Array := strings.Fields(row3)
	row4Array := strings.Fields(row4)
	if width == 5 {
		row5Array = strings.Fields(row5)
	}

	dice[0][0].character = row1Array[0]
	dice[0][0].id = 0
	dice[0][1].character = row1Array[1]
	dice[0][1].id = 1
	dice[0][2].character = row1Array[2]
	dice[0][2].id = 2
	dice[0][3].character = row1Array[3]
	dice[0][3].id = 3
	if width == 5 {
		dice[0][4].character = row1Array[4]
		dice[0][4].id = 4
	}

	dice[1][0].character = row2Array[0]
	dice[1][0].id = 5
	dice[1][1].character = row2Array[1]
	dice[1][1].id = 6
	dice[1][2].character = row2Array[2]
	dice[1][2].id = 7
	dice[1][3].character = row2Array[3]
	dice[1][3].id = 8
	if width == 5 {
		dice[1][4].character = row2Array[4]
		dice[1][4].id = 9
	}

	dice[2][0].character = row3Array[0]
	dice[2][0].id = 10
	dice[2][1].character = row3Array[1]
	dice[2][1].id = 11
	dice[2][2].character = row3Array[2]
	dice[2][2].id = 12
	dice[2][3].character = row3Array[3]
	dice[2][3].id = 13
	if width == 5 {
		dice[2][4].character = row3Array[4]
		dice[2][4].id = 14
	}

	dice[3][0].character = row4Array[0]
	dice[3][0].id = 15
	dice[3][1].character = row4Array[1]
	dice[3][1].id = 16
	dice[3][2].character = row4Array[2]
	dice[3][2].id = 17
	dice[3][3].character = row4Array[3]
	dice[3][3].id = 18
	if width == 5 {
		dice[3][4].character = row4Array[4]
		dice[3][4].id = 19
	}

	if width == 5 {
		dice[4][0].character = row5Array[0]
		dice[4][0].id = 20
		dice[4][1].character = row5Array[1]
		dice[4][1].id = 21
		dice[4][2].character = row5Array[2]
		dice[4][2].id = 22
		dice[4][3].character = row5Array[3]
		dice[4][3].id = 23
		dice[4][4].character = row5Array[4]
		dice[4][4].id = 24
	}

	return dice
}

func PrintDice(dice [5][5]diceValue) {
	if width == 4 {
		fmt.Println(dice[0][0].character, "\t", dice[0][1].character, "\t", dice[0][2].character, "\t", dice[0][3].character)
		fmt.Println(dice[1][0].character, "\t", dice[1][1].character, "\t", dice[1][2].character, "\t", dice[1][3].character)
		fmt.Println(dice[2][0].character, "\t", dice[2][1].character, "\t", dice[2][2].character, "\t", dice[2][3].character)
		fmt.Println(dice[3][0].character, "\t", dice[3][1].character, "\t", dice[3][2].character, "\t", dice[3][3].character)
	} else if width == 5 {
		fmt.Println(dice[0][0].character, "\t", dice[0][1].character, "\t", dice[0][2].character, "\t", dice[0][3].character, "\t", dice[0][4].character)
		fmt.Println(dice[1][0].character, "\t", dice[1][1].character, "\t", dice[1][2].character, "\t", dice[1][3].character, "\t", dice[1][4].character)
		fmt.Println(dice[2][0].character, "\t", dice[2][1].character, "\t", dice[2][2].character, "\t", dice[2][3].character, "\t", dice[2][4].character)
		fmt.Println(dice[3][0].character, "\t", dice[3][1].character, "\t", dice[3][2].character, "\t", dice[3][3].character, "\t", dice[3][4].character)
		fmt.Println(dice[4][0].character, "\t", dice[4][1].character, "\t", dice[4][2].character, "\t", dice[4][3].character, "\t", dice[4][4].character)
	}
}
