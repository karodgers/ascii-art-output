package handlers

import (
	"fmt"
	"os"
	"strings"
)

// ExtractTextAndFilename extracts the text and filename from command-line arguments.
func ExtractTextAndFilename(args []string) (string, string) {
	text := ""
	filename := ""

	// Process the first argument if it starts with --output=
	if strings.HasPrefix(args[0], "--output=") {
		filename = strings.TrimPrefix(args[0], "--output=")
		if !strings.HasSuffix(filename, ".txt") {
			fmt.Println("Error: Must have a .txt file for printing output.")
			os.Exit(0)
		}
	}

	// The last argument should be the banner type
	lastArg := args[len(args)-1]
	switch lastArg {
	case "thinkertoy", "shadow", "standard":
		text = extractText(args[:len(args)-1])
	default:
		if !strings.HasPrefix(args[0], "--output=") {
			PrintUsage()
			os.Exit(0)
		} else {
			fmt.Println("Error: Please specify a valid banner type. Use go run . [OPTION] [STRING] [Standard/Thinkertoy/Shadow]")
			os.Exit(0)
		}

	}

	return text, filename
}

// extractText joins arguments into a single string excluding any --output= flag.
func extractText(args []string) string {
	str := strings.Join(args, " ")
	// Remove any --output= flag part from the text
	for _, arg := range args {
		if strings.HasPrefix(arg, "--output=") {
			str = strings.ReplaceAll(str, arg, "")
		}
	}
	return strings.TrimSpace(str)
}

// PrintUsage prints the usage message.
func PrintUsage() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
	fmt.Println("Example: go run . --output=banner.txt something standard")
}
