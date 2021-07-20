package main

import (
	"cyancoding/go-humanize"
	"fmt"
	"io/ioutil"
	"strings"
	"sync"

	"github.com/slok/gospinner"
)

// dictionary.txt contains words you can find in a dictionary (valid boggle words)
// dictionary-alt.txt is the same as dictionary.txt with a few capitalized things
// like weekdays, a few common names, a few abbreviations, etc.
// all-english.txt contains every english word (might not be in a dictionary)
var wordsFile string

var wordsFound int
var searches int64
var wordsChecked int64
var words []string // Word list (about 436k)
var wordsFoundList []string

var s *gospinner.Spinner

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
	dice := RollDice()

	fmt.Println("Welcome to the boggle solver/computer version!")
	fmt.Println("Please choose one of the following")
	fmt.Println("Solo mode (1), computer solve (2), input board (3)")
	fmt.Println()
	fmt.Print("Action > ")

	var action int = 2
	fmt.Scanln(&action)

	fmt.Println()
	fmt.Println("Please pick your dictionary.")
	fmt.Println("(1) Only dictionary words, (2) all English words")
	fmt.Print("Dictionary > ")
	var dictionary int = 0
	fmt.Scanln(&dictionary)
	fmt.Println()

	if dictionary == 1 {
		wordsFile = "dictionary.txt"
	} else if dictionary == 2 {
		wordsFile = "all-english.txt"
	} else {
		fmt.Println("Invalid response! Defaulting to dictionary.")
		wordsFile = "dictionary.txt"
	}

	var group sync.WaitGroup

	group.Add(1)
	go func() {
		words = ReadWordsFile()
		defer group.Done()
	}()

	if action < 1 || action > 3 { // Action is not valid
		action = 2
		fmt.Println("\nYou did not enter a valid number. Defaulting to 2")
	}

	group.Wait() // Make sure the word compilation has finished
	wordsFoundList = make([]string, 0)

	if action == 3 {
		dice = ManuallySetDice()
	} else if action == 1 {
		fmt.Println("This feature is not ready yet!")
	}

	fmt.Println("Below is the board")
	fmt.Println()
	PrintDice(dice)
	fmt.Println()
	s, _ = gospinner.NewSpinner(gospinner.Dots2)
	s.Start("Finding matches (0)...")

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			FindNearby(dice, i, j, "")
		}
	}
	s.SetMessage("Found all matches!")
	s.Succeed()

	for _, a := range wordsFoundList {
		fmt.Println(a)
	}

	fmt.Println("Ran", humanize.Comma(searches), "times and tested",
		humanize.Comma(wordsChecked), "words")

	var letterCount int64
	for _, a := range wordsFoundList {
		letterCount += int64(len(a))
	}

	fmt.Println("That's", humanize.Comma(int64(len(wordsFoundList))), "words and",
		humanize.Comma(letterCount), "points!")
}
