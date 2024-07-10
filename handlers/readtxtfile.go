package handlers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	asciiArtHeight = 8
)

// used to read the .txt asciiArt file which are then stored in the variable asciiArt as a slice of string according to their indexes
// defines the banner file flags to be used for modifying the output
func ReadAsciiArt() []string {
	var args []string

	args = append(args, os.Args[1:]...)

	var filename string

	switch args[len(args)-1] {
	case "thinkertoy":
		filename = "thinkertoy.txt"
	case "shadow":
		filename = "shadow.txt"
	case "standard":
		filename = "standard.txt"
	default:
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error while opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	lineCount := 0

	scanner := bufio.NewScanner(file)

	var asciiArt []string
	var artLines []string

	for scanner.Scan() {

		lineCount++

		lines := scanner.Text()

		if len(lines) == 0 {
			continue
		}
		artLines = append(artLines, lines)

		if len(artLines) == asciiArtHeight {
			asciiArt = append(asciiArt, strings.Join(artLines, "\n"))
			artLines = nil
		}
	}

	if scanner.Err() != nil {
		fmt.Println("Error reading ASCII art file:", scanner.Err())
		os.Exit(1)
	}
	if len(asciiArt) == 0 {
		fmt.Println("Error: The ASCII art file is empty.")
		os.Exit(1)
	}

	if lineCount != 855 {
		fmt.Println("Read Error: Do not change the content of the txt file")
		os.Exit(1)
	}

	return asciiArt
}
