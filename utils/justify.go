package utils

import (
	"fmt"
	"os"
	"strings"
	"syscall"
	"unsafe"
)

func Justify(flagValue string, text string, asciiArt map[rune][]string) {
	var layers [8]string
	var art string
	var inputValue []string

	width, _, err := getTerminalSize()
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		return
	}

	lines := strings.Split(text, `\n`)
	if flagValue == "justify" {
		var extraSpaces int
		for _, line := range lines {
			if line == "" {
				art = fmt.Sprint(art, "\n")
			} else {
				rows := [8][]string{}
				rowLength := 0
				spaceCount := 0
				for i := 0; i < 8; i++ {
					str := ""
					for _, char := range line {
						if char == ' ' {
							rows[i] = append(rows[i], str)
							str = ""
							if i == 0 {
								spaceCount++
							}
						} else {
							str = fmt.Sprint(str, asciiArt[char][i])
							if i == 0 {
								rowLength += len(asciiArt[char][i])
							}
						}
					}
					rows[i] = append(rows[i], str)
				}
				if rowLength > width {
					fmt.Println("Text is too long for terminal width!")
					os.Exit(1)
				}
				totalSpaces := width - rowLength - spaceCount
				spacesBetweenWords := 1
				if spaceCount > 0 {
					spacesBetweenWords += totalSpaces / spaceCount
				}
				if spaceCount != 0 {
					extraSpaces = totalSpaces % spaceCount
				}

				for i := range rows {
					row := ""
					for j, word := range rows[i] {
						row += word
						if j < len(rows[i])-1 {
							spacesToAdd := spacesBetweenWords
							if extraSpaces > 0 {
								extraSpaces--
							}
							row += strings.Repeat(" ", spacesToAdd)
						}
					}
					fmt.Println(row)
				}
			}
		}
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
	if len(layer) > width {
		fmt.Println("Text is too long for terminal width!")
		os.Exit(1)
	}
	switch alignment {
	case "left":
		fmt.Println(layer)
	case "center":
		padding := strings.Repeat(" ", (width-len(layer))/2)
		fmt.Println(padding + layer)
	case "right":
		padding := strings.Repeat(" ", width-len(layer))
		fmt.Println(padding + layer)
	}
}
