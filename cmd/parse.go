package cmd

import "strings"

// parsing input
func ParseInput(input string) []string {
	return strings.Fields(input)
}
