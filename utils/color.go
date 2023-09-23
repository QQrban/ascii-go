package utils

import (
	"fmt"
	"os"
	"strings"
)

// ...

func Color(textToColor string, asciiArt map[rune][]string, selectedColor string, charsToColor []string) {
	var layers [8]string

	wordsToColor := make(map[string]bool)
	for _, s := range charsToColor {
		wordsToColor[s] = true
	}

	lines := strings.Split(textToColor, "\\n")

	for _, line := range lines {
		if line == "" {
			fmt.Println()
			continue
		}
		words := strings.Fields(line)
		for _, word := range words {
			wordColored := false

			if wordsToColor[word] {
				for _, char := range word {
					if art, exists := asciiArt[char]; exists {
						for j, line := range art {
							layers[j] += selectedColor + line + " " + colors["reset"]
						}
					} else {
						fmt.Println("One or more symbols don't exist in ASCII-table!")
						os.Exit(1)
					}
				}
				wordColored = true
			}

			if !wordColored {
				for _, char := range word {
					charColored := false
					if wordsToColor[string(char)] {
						charColored = true
						if art, exists := asciiArt[char]; exists {
							for j, line := range art {
								layers[j] += selectedColor + line + " " + colors["reset"]
							}
						} else {
							fmt.Println("One or more symbols don't exist in ASCII-table!")
							os.Exit(1)
						}
					}

					if !charColored {
						if art, exists := asciiArt[char]; exists {
							for j, line := range art {
								layers[j] += line + " "
							}
						} else {
							fmt.Println("One or more symbols don't exist in ASCII-table!")
							os.Exit(1)
						}
					}
				}
			}

			for j := range layers {
				layers[j] += " "
			}
		}

		if textToColor != "" && len(charsToColor) == 0 {
			for i := 0; i < 8; i++ {
				fmt.Println(selectedColor + layers[i] + colors["reset"])
			}
			fmt.Println()
			return
		}

		// После обработки каждой строки выводим содержимое слоев и очищаем их для следующей строки
		for _, layer := range layers {
			fmt.Println(layer + colors["reset"])
		}
		layers = [8]string{}
	}
	fmt.Println()
}

func checkColor(flagValue string) string {
	selectedColor := ""
	if _, exists := colors[flagValue]; exists {
		selectedColor = colors[flagValue]
	} else {
		fmt.Println("Sorry, there is no selected color :(")
		os.Exit(1)
	}

	return selectedColor
}

func handleErrors(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: go run . --color=<color> <text>")
		os.Exit(1)
	}

	textArgs := args[2:]
	if len(textArgs) == 0 {
		fmt.Println("You must provide text to print")
		os.Exit(1)
	}

	var prefixes = []string{"--output", "--color", "--reverse"}

	for _, prefix := range prefixes {
		if strings.HasPrefix(args[2], prefix) {
			fmt.Println("Error: Only 1 flag can be provided!")
			os.Exit(1)
		}
	}
}

var colors = map[string]string{
	"red":         "\033[38;2;255;0;0m",
	"green":       "\033[38;2;0;255;0m",
	"blue":        "\033[38;2;0;0;255m",
	"yellow":      "\033[38;2;255;255;0m",
	"purple":      "\033[38;2;255;0;255m",
	"cyan":        "\033[38;2;0;255;255m",
	"orange":      "\033[38;2;255;165;0m",
	"violet":      "\033[38;2;128;0;128m",
	"teal":        "\033[38;2;0;128;128m",
	"springGreen": "\033[38;2;0;255;127m",
	"gold":        "\033[38;2;255;215;0m",
	"crimson":     "\033[38;2;220;20;60m",
	"darkGreen":   "\033[38;2;0;128;0m",
	"darkBlue":    "\033[38;2;0;0;128m",
	"brown":       "\033[38;2;165;42;42m",
	"reset":       "\033[0m",
}
