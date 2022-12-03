package AocHelper

import (
	"fmt"
	"os"
)

func cacheInput(rawInput string, sessionID string, year int, day int) error {
	err := os.WriteFile(fmt.Sprintf("%s/%s_%d_%d.txt", os.TempDir(), sessionID, year, day), []byte(rawInput), 0644)
	if err != nil {
		return err
	}

	return nil
}

func checkForCachedInput(sessionID string, year int, day int) bool {
	if _, err := os.Stat(fmt.Sprintf("%s/%s_%d_%d.txt", os.TempDir(), sessionID, year, day)); !os.IsNotExist(err) {
		return true
	}
	
	return false
}

func readCachedInput(sessionID string, year int, day int) (input, error) {
	var in input 

	data, err := os.ReadFile(fmt.Sprintf("%s/%s_%d_%d.txt", os.TempDir(), sessionID, year, day))
	if err != nil {
		return in, err
	}

	in.rawInput = string(data)

	return in, nil
}
