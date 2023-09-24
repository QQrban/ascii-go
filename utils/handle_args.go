package utils

import (
	"regexp"
	"strings"
)

var flagPattern = regexp.MustCompile(`^--(output|reverse|color|align)=([^ ]+)$`)

func HandleArgs(args string) (string, string) {
	matches := flagPattern.FindStringSubmatch(args)
	if len(matches) > 2 {
		return matches[1], matches[2]
	}
	return "", ""
}

func IsFlagged(args string) bool {
	return strings.HasPrefix(args, "--")
}
