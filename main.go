package main

import (
	"ascii-art/utils"
	"fmt"
	"os"
)

func main() {
	// Load ASCII art map.

	// Check if there are command-line arguments.
	if len(os.Args) < 2 {
		fmt.Println("Please enter your text")
		return
	}

	utils.WriteFile(os.Args)
}
