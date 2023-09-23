package utils

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// Reverse function reads an ASCII art from a file and converts it back to its original text representation.
// It uses a predefined map of symbols to ASCII art (`symbolMap`) to find the corresponding text character for each ASCII art block.
func Reverse(fileName string, symbolMap map[rune][]string) string {

	// Read the content of the specified file.

	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading the file:", err)
		os.Exit(1)
	}

	// Split the file content into individual lines.
	lines := strings.Split(string(fileContent), "\n")

	// Ensure the file has at least 9 lines (8 for ASCII art and 1 empty).
	if len(lines) < 9 {
		fmt.Println("File content is too short.")
		return ""
	}

	result := []rune{}
	sortedPairs := sortedSymbolMap(symbolMap)

	// Process the file as long as there's content left.
	for len(lines[0]) > 0 {
		found := false

		// Try to match each ASCII art block in the file to a character in the symbol map.
		for _, pair := range sortedPairs { // Итерируемся по отсортированному списку
			key := pair.Key
			value := pair.Value
			match := true
			for i := 0; i < 8; i++ {
				if !strings.HasPrefix(lines[i], value[i]) {
					match = false
					break
				}
			}
			if match {
				result = append(result, key)
				found = true
				for i := 0; i < 8; i++ {
					lines[i] = lines[i][len(value[i]):]
					if len(lines[i]) > 0 {
						lines[i] = lines[i][1:]
					}
				}
				break
			}
		}
		if found {
			allEmpty := true
			for i := 0; i < 8; i++ {
				if len(lines[i]) > 0 {
					allEmpty = false
					break
				}
			}
			if allEmpty {
				result = append(result, '\n')
				lines = lines[8:]
			}
		} else {
			fmt.Println("Error: Unable to find a match for the ASCII art.")
			os.Exit(1)
		}
	}
	return string(result)
}

type Pair struct {
	Key   rune
	Value []string
}

func sortedSymbolMap(symbolMap map[rune][]string) []Pair {
	pairs := make([]Pair, 0, len(symbolMap))
	for k, v := range symbolMap {
		pairs = append(pairs, Pair{k, v})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return len(pairs[i].Value[0]) > len(pairs[j].Value[0])
	})
	return pairs
}
