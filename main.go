package main

import (
	"fmt"
	"os"
	"strings"

	"ascii/handlers"
)

// calls back the logic functions for printing the output to the banner file
func main() {
	args := os.Args[1:]
	if len(args) == 0 || len(args) > 3 {
		handlers.PrintUsage()
		os.Exit(0)
	}

	text, filename := handlers.ExtractTextAndFilename(args)
	if text == "" {
		handlers.PrintUsage()
		os.Exit(0)
	}
	asciiArt := handlers.ReadAsciiArt()

	if handlers.ContainsNonASCII(text) {
		fmt.Println("Print Error: Non ASCII characters detected")
		os.Exit(0)
	}

	contents := handlers.PrintAsciiArt(text, asciiArt)

	if !strings.HasPrefix(args[0], "--output=") {
		handlers.PrintUsage()
		os.Exit(0)
	}

	err := os.WriteFile(filename, []byte(contents), 0o664)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

}
