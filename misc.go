package AocHelper

import "strings"

func getLines(rawInput string) []string {
	var lines []string 

	for _, v := range strings.Split(rawInput, "\n") {
		lines = append(lines, v)
	}
	
	lines = lines[:len(lines) - 1]
	
	return lines
}
