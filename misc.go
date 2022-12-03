package AocHelper

import "strings"

func getLines(rawInput string) []string {
	var lines []string 

	for _, v := range strings.Split(rawInput, "\n") {
		if v != "" {
			lines = append(lines, v)
		}
	}
	
	return lines
}
