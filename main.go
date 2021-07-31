package main

import (
	"cyancoding/go-humanize"
	"fmt"
	"io/ioutil"
	"strings"
	"sync"

	"github.com/slok/gospinner"
)

// Colors
const ColorReset = "\033[0m"
const ColorRed = "\033[31m"
const ColorGreen = "\033[32m"
const ColorYellow = "\033[33m"
const ColorBlue = "\033[34m"
const ColorPurple = "\033[35m"
const ColorCyan = "\033[36m"

// oxford.txt contains words you can find in a dictionary (valid boggle words)
// all-english.txt contains every english word (might not be in a dictionary)
var wordsFile string = "oxford.txt"

var wordsFound int
var searches int64
var wordsChecked int64
var words []string // Word list (about 436k)
var wordsFoundList []string

var s *gospinner.Spinner

var percentDone int

var width int = 4 // 4 or 5 for 4x4/5x5

func ReadWordsFile() []string {
	byteData, err := ioutil.ReadFile(wordsFile)

	if err != nil {
		fmt.Println(err)
	}

	text := string(byteData)
	words := strings.Fields(text)

	return words
}

func main() {
	fmt.Println(ColorPurple + "Welcome to the boggle solver/computer version!")
	fmt.Println("Please choose one of the following")
	fmt.Println("Solo mode (1), computer solve (2), input board (3)")
	fmt.Println()
	fmt.Print(ColorYellow + "Action > ")

	var action int = 2
	fmt.Scanln(&action)

	fmt.Println()
	fmt.Println(ColorPurple + "Please pick a board size")
	fmt.Print(ColorYellow + "(1) 4x4, (2) 5x5 > ")

	var boardSize int = 1
	fmt.Scanln(&boardSize)

	if boardSize == 2 {
		width = 5
	} else if boardSize != 1 { // Invalid value
		fmt.Println(ColorRed + "Invalid response! Defaulting to 4x4")
		width = 4
	}

	dice := RollDice()

	fmt.Println()
	fmt.Println(ColorPurple + "Please pick your dictionary")
	fmt.Println("(1) Only dictionary words, (2) all English words")
	fmt.Print(ColorYellow + "Dictionary > ")
	var dictionary int = 1
	fmt.Scanln(&dictionary)
	fmt.Println()

	if dictionary == 2 {
		wordsFile = "all-english.txt"
	} else if dictionary != 1 {
		fmt.Println(ColorRed + "Invalid response! Defaulting to dictionary")
	}

	fmt.Println()
	fmt.Println(ColorPurple + "Please pick an option")
	fmt.Println("(1) Use repeat letters, (2) no repeat letters")
	fmt.Print(ColorYellow + "Option > ")
	var option int = 0
	var useRepeats bool = true
	fmt.Scanln(&option)
	fmt.Println()

	if option == 2 {
		useRepeats = false
	} else if option != 1 {
		fmt.Println(ColorRed + "Invalid response! Defaulting to yes for repeat letters")
	}

	var group sync.WaitGroup

	group.Add(1)
	go func() {
		words = ReadWordsFile()
		defer group.Done()
	}()

	if action < 1 || action > 3 { // Action is not valid
		action = 2
		fmt.Println(ColorRed + "\nYou did not enter a valid number. Defaulting to 2")
	}

	group.Wait() // Make sure the word compilation has finished
	wordsFoundList = make([]string, 0)

	if action == 3 {
		dice = ManuallySetDice()
		fmt.Println(ColorRed + "This feature is not ready yet!")
	} else if action == 1 {
		fmt.Println(ColorRed + "This feature is not ready yet!")
	}

	fmt.Println(ColorPurple + "Below is the board")
	fmt.Println()
	PrintDice(dice)
	fmt.Println()
	s, _ = gospinner.NewSpinner(gospinner.Dots2)
	s.Start(ColorCyan + "Finding matches (0% • 0)...")

	for i := 0; i < width; i++ {
		for j := 0; j < width; j++ {
			FindNearby(dice, i, j, "", true, nil)
			percentDone += 100 / (width * width)
		}
	}
	s.SetMessage(ColorGreen + "Found all matches!")
	s.Succeed()

	for _, a := range wordsFoundList {
		fmt.Println(a)
	}

	fmt.Println(ColorPurple+"Ran", humanize.Comma(searches), "times and tested",
		humanize.Comma(wordsChecked), "words")

	var letterCount int64
	for _, a := range wordsFoundList {
		letterCount += int64(len(a))
	}

	fmt.Println("That's", humanize.Comma(int64(len(wordsFoundList))), "words and",
		humanize.Comma(letterCount), "points!")
}
