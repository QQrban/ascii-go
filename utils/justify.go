package utils

import (
	"fmt"
	"os"
	"strings"
	"syscall"
	"unsafe"
)

func Justify(flagValue string, text string, asciiArt map[rune][]string) {
	var (
		inputValue []string
		layers     [8]string
	)

	width, _, err := getTerminalSize()
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		return
	}

	for _, s := range text {
		inputValue = append(inputValue, string(s))
	}

	for i := 0; i < len(inputValue); i++ {

		if inputValue[i] == "\\" && i < len(inputValue)-1 && inputValue[i+1] == "n" {
			// Выводим текущие слои и сбрасываем их для новой строки
			for _, layer := range layers {
				printAlignedText(layer, flagValue, width)
			}

			layers = [8]string{}
			i++

			for i+1 < len(inputValue) && inputValue[i+1] == "\\" && i+2 < len(inputValue) && inputValue[i+2] == "n" {
				fmt.Println()
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
		printAlignedText(layer, flagValue, width)
	}
	fmt.Println()

}

func getTerminalSize() (width int, height int, err error) {
	var ws struct {
		rows    uint16
		cols    uint16
		xpixels uint16
		ypixels uint16
	}
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(os.Stdout.Fd()), uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&ws)))
	if errno != 0 {
		err = errno
		return
	}
	width = int(ws.cols)
	height = int(ws.rows)
	return
}

func printAlignedText(layer string, alignment string, width int) {
	spaces := 0
	padding := ""
	switch alignment {
	case "left":
	case "center":
		spaces = (width - len(layer)) / 2
		padding = strings.Repeat(" ", spaces)
	case "right":
		spaces = width - len(layer)
		padding = strings.Repeat(" ", spaces)
	case "justify":
	}
	fmt.Println(padding + layer)
}
