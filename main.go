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
const (
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"

	version = 1.3
)

// wordsFile is the relative path to the dictionary in use.
// oxford.txt contains words you can find in a dictionary (valid boggle words)
// all-english.txt contains every english word, like company names (might not be in a dictionary)
var wordsFile string = "oxford.txt"

var wordsFound int
var searches int64
var wordsChecked int64
var words []string // Word list (about 38K items depending on the dictionary)
var wordsFoundList []string

// Used for the progress bar
var s *gospinner.Spinner
var percentDone int

var width int = 4 // 4 or 5 for 4x4/5x5

// readWordsFile takes wordsFile and reads line by line into words[]
func readWordsFile() []string {
	byteData, err := ioutil.ReadFile(wordsFile)

	if err != nil {
		fmt.Println(err)
	}

	text := string(byteData)
	words := strings.Fields(text)

	return words
}

// pickDictionary prompts the user and fills in wordsFile
func pickDictionary() {
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
		fmt.Println(ColorRed + "Invalid response! Defaulting to oxford dictionary")
	}
}

func main() {
	fmt.Println(ColorPurple + "Welcome to the boggle solver/computer version by CyanCoding!")
	fmt.Println("Version:", version)
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

	pickDictionary()

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

	var dice [5][5]diceValue

	group.Add(1)
	go func() {
		words = readWordsFile()
		if action != 3 {
			dice = RollDice()
		}
		defer group.Done()
	}()

	if action < 1 || action > 3 { // Action is not valid
		action = 2
		fmt.Println(ColorRed + "\nYou did not enter a valid number. Defaulting to 2")
	}

	if action == 3 {
		dice = ManuallySetDice()
	} else if action == 1 {
		fmt.Println(ColorRed + "This feature is not ready yet!")
	}

	group.Wait() // Make sure the word compilation has finished
	wordsFoundList = make([]string, 0)

	fmt.Println(ColorPurple + "Below is the board")
	fmt.Println()
	PrintDice(dice)
	fmt.Println()
	s, _ = gospinner.NewSpinner(gospinner.Dots2)
	s.Start(ColorCyan + "Finding matches (0% • 0)...")

	// This is the actual for loop that runs the program
	for i := 0; i < width; i++ {
		for j := 0; j < width; j++ {
			FindNearby(dice, i, j, "", useRepeats, nil)
			percentDone += 100 / (width * width)
		}
	}
	s.SetMessage(ColorGreen + "Found all matches!")
	s.Succeed()

	// Reports all of the words when we're done
	for _, a := range wordsFoundList {
		fmt.Println(a)
	}

	// Count the letters
	var letterCount int64
	for _, a := range wordsFoundList {
		letterCount += int64(len(a))
	}

	// Give the user statistics
	fmt.Println(ColorPurple+"Ran", humanize.Comma(searches), "times and tested",
		humanize.Comma(wordsChecked), "words")
	fmt.Println("That's", humanize.Comma(int64(len(wordsFoundList))), "words and",
		humanize.Comma(letterCount), "points!")
}
