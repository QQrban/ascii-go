package utils

import (
	"bufio"
	"os"
)

func LoadMap(style string) map[rune][]string {

	file, err := os.Open(style)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	// Initialize the map to store ASCII art for each character.
	asciiMap := make(map[rune][]string)

	// Iterate through ASCII characters from 32 to 126.
	for i := 32; i <= 126; i++ {
		var lines []string

		// Read 8 lines of ASCII art for each character.
		for j := 0; j < 8; j++ {
			if scanner.Scan() {
				lines = append(lines, scanner.Text())
			}
		}

		// Store the ASCII art lines for the current character.
		asciiMap[rune(i)] = lines

		// Skip the empty line between characters.
		scanner.Scan()
	}
	return asciiMap
}
