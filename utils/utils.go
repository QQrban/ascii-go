package utils

import (
	"fmt"
	"regexp"
	"strings"
)

// WriteFile generates ASCII art from the given input text and writes it to a specified file.
func WriteFile(args []string) {

	var text string

	var style string

	if strings.HasSuffix(args[len(args)-1], ".txt") {
		style = "standard"
	} else {
		style = args[len(args)-1]
	}
	selectedStyle := GetStyle(style)

	// Load the ASCII art characters from the selected style file.
	asciiArt := LoadMap(selectedStyle + ".txt")
	// Determine the start and end indices for processing command-line arguments.
	startIndex, endIndex := processArgs(args)
	// Extract the input text from the command-line arguments.
	text = strings.Join(args[startIndex:endIndex], " ")
	_, err := checkString(text)
	if err != nil {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("EX: go run . something standard")
		return
	}
	flagType, flagValue := HandleArgs(args[1])
	if IsFlagged(args[1]) {
		if (flagType == "" || flagValue == "") && flagType != "color" {
			fmt.Println("Error: Invalid flag. Use --output=<filename>.txt, --reverse=<filename>.txt")
			fmt.Println("or --color=<color>")
			return
		}
	}

	// Convert the input text into individual characters for ASCII conversion.
	switch flagType {
	case "output":
		Output(text, asciiArt, flagValue, 1)
	case "reverse":
		reversedString := Reverse(flagValue, asciiArt)
		fmt.Println(reversedString)
	case "color":
		handleErrors(args)
		var charsToColor []string
		var textToColor string
		styles := []string{"standard", "shadow", "thinkertoy"}
		if len(args[1:]) > 2 {
			charsToColor = args[2 : len(args)-1]
		}
		textToColor = args[len(args)-1]
		for _, prefix := range styles {
			if strings.HasPrefix(args[len(args)-1], prefix) {
				textToColor = args[len(args)-2]
				break
			}
		}
		selectedColor := checkColor(flagValue)

		Color(textToColor, asciiArt, selectedColor, charsToColor)
	case "":
		Output(text, asciiArt, flagValue, 2)
	}
}

func processArgs(args []string) (int, int) {
	stylePattern := regexp.MustCompile("(shadow|standard|thinkertoy)")

	var startIndex, endIndex int

	if IsFlagged(args[1]) {
		startIndex = 2
	} else {
		startIndex = 1
	}

	if stylePattern.MatchString(args[len(args)-1]) {
		endIndex = len(args) - 1
	} else {
		endIndex = len(args)
	}

	return startIndex, endIndex
}

func checkString(s string) (string, error) {
	parts := strings.Fields(s)
	styles := []string{"standard", "shadow", "thinkertoy"}

	for i, part := range parts {
		for _, style := range styles {
			if part == style {
				if i < len(parts)-1 {
					return "", fmt.Errorf("error: Invalid arguments after style")
				}
			}
		}
	}
	return s, nil
}
