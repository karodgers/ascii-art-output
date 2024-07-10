package handlers

import (
	"fmt"
	"strings"
)

// This function pre-checks different conditions passed as arguments, including (empty strings, special charachters, and string literals), and handles errors where necessary.
func PrintAsciiArt(text string, asciiArt []string) string {
	var builder1 strings.Builder
	if text == "" {
		return ""
	}
	// error handling for special charachters
	specialChars := map[string]string{
		"\\t": "Tab",
		"\\b": "Backspace",
		"\\v": "Vertical Tab",
		"\\0": "Null",
		"\\f": "Form Feed",
		"\\r": "Carriage Return",
	}

	for spChar, description := range specialChars {
		if strings.Contains(text, spChar) {
			fmt.Printf("Print Error: Special ASCII character (%s) or (%s) detected \n", spChar, description)
			return ""

		}
	}

	// handling cases of string literals
	if strings.Contains(text, "\\n") {
		input := strings.Split(text, "\\n")
		count := 0
		for _, word := range input {
			if word == "" {
				count++
				if count < len(input) {
					fmt.Println()
				}
			} else if len(word) > 0 {
				_, err := builder1.WriteString(PrintLineByLine(word, asciiArt))
				if err != nil {
					return ""
				}
			}
		}
	} else {

		lines := strings.Split(text, "\n")
		for _, line := range lines {
			if len(line) > 0 {
				_, err := builder1.WriteString(PrintLineByLine(line, asciiArt))
				if err != nil {
					return ""
				}
			} else {
				fmt.Println()
			}
		}
	}
	return builder1.String()
}
