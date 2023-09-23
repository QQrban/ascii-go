package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func Output(text string, asciiArt map[rune][]string, flagValue string, variant int) {
	var (
		file       *os.File
		err        error
		inputValue []string
		layers     [8]string
	)

	for _, s := range text {
		inputValue = append(inputValue, string(s))
	}

	if variant == 1 {
		file, err = prepareOutputFile(flagValue)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()
	}

	// Iterate through the input characters.
	for i := 0; i < len(inputValue); i++ {

		// Handle '\n' character, if found.
		if inputValue[i] == "\\" && i < len(inputValue)-1 && inputValue[i+1] == "n" {
			// Print the existing layers, then add an empty line and reset current layers
			for _, layer := range layers {
				if variant == 1 {
					_, writeErr := file.WriteString(layer + "\n")
					if writeErr != nil {
						fmt.Println("Error writing to file:", writeErr)
						os.Exit(1)
					}
				} else if variant == 2 {
					fmt.Println(layer)
				}
			}
			layers = [8]string{}
			i++

			for i+1 < len(inputValue) && inputValue[i+1] == "\\" && i+2 < len(inputValue) && inputValue[i+2] == "n" {
				if variant == 1 {
					_, writeErr := file.WriteString("\n")
					if writeErr != nil {
						fmt.Println("Error writing to file:", writeErr)
						os.Exit(1)
					}
				} else if variant == 2 {
					fmt.Println()
				}
				i += 2
			}

			continue
		}

		if art, exists := asciiArt[rune(inputValue[i][0])]; exists {
			for j, line := range art {
				layers[j] += line + " "
			}
		} else {
			fmt.Println("One or more symbols don't exist in ASCII-table!")
			os.Exit(1)
		}
	}

	for _, layer := range layers {
		if variant == 1 {
			_, writeErr := file.WriteString(layer + "\n")
			if writeErr != nil {
				fmt.Println("Error writing to file:", writeErr)
				os.Exit(1)
			}
		} else if variant == 2 {
			fmt.Println(layer)
		}
	}

	if variant == 1 {
		fmt.Println("File created!")
	} else if variant == 2 {
		fmt.Println()
	}
}

func prepareOutputFile(fileName string) (*os.File, error) {
	if filepath.Ext(fileName) != ".txt" {
		return nil, fmt.Errorf("OOPS! File extension must be .txt. (i.e. test.txt)  :)")
	}

	file, err := os.Create(fileName)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	return file, nil
}
