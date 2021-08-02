package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"
	"unicode"

	"github.com/go-vgo/robotgo"
	"github.com/otiai10/gosseract/v2"
)

func colorEqual(c1, c2 color.Color) bool {
	r1, g1, b1, a1 := c1.RGBA()
	r2, g2, b2, a2 := c2.RGBA()

	return r1 == r2 && g1 == g2 && b1 == b2 && a1 == a2
}

func GetText() string {
	// You can register another format here
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	fmt.Println("Taking screenshot...")

	bitmap2 := robotgo.CaptureScreen()
	defer robotgo.FreeBitmap(bitmap2)
	robotgo.SaveBitmap(bitmap2, "screenshot.png")

	fmt.Println("Converting screenshot...")
	file, _ := os.Open("screenshot.png")
	defer file.Close()

	img, _, _ := image.Decode(file)

	bounds := img.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y

	i := 0
	matchX := 0
	matchY := 0
	foundMatch := false

	fmt.Println("Finding color matches...")

	for y := 0; y < height; y++ {
		if foundMatch {
			break
		}

		for x := 0; x < width; x++ {
			c := img.At(x, y)

			if colorEqual(c, color.RGBA{223, 160, 100, 255}) {
				i++
				if i == 1 {
					matchX = x
					matchY = y

					fmt.Println("Match at", x, y)
					foundMatch = true
					break
				}
			}
		}
	}

	if !foundMatch {
		fmt.Println("No matches! :(")
		return ""
	}

	bitmap := robotgo.CaptureScreen(matchX+150, matchY, 340, 350)
	defer robotgo.FreeBitmap(bitmap)

	robotgo.SaveBitmap(bitmap, "board.png")

	fmt.Println("Reading text...")

	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage("board.png")
	text, _ := client.Text()
	newText := ""

	//fmt.Println(text)

	i = 0

	fmt.Println("Converting text...")

	for _, c := range text {
		if unicode.IsLetter(c) && unicode.IsUpper(c) {
			i++
			newText += string(c) + " "

			if i%4 == 0 {
				newText += "\n"
			}
		}
	}

	return strings.ToLower(newText)
}
