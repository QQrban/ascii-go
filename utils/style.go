package utils

import "regexp"

func GetStyle(style string) string {
	stylePattern := regexp.MustCompile("(shadow|standard|thinkertoy)")

	match := stylePattern.FindString(style)
	if match != "" {
		return match
	}
	return "standard"
}
