package AocHelper

import (
	"io/ioutil"
	"net/http"
	"fmt"
	"os"
)

func (session *Session) InitSession(sessionID string) {
	session.sessionID = sessionID
}

func (session *Session) InitSessionFromFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	session.InitSession(string(data))

	return nil
}

func (session Session) RetrieveInput(year int, day int) (input, error) {
	var in input

	hasCachedInput := checkForCachedInput(session.sessionID, year, day)
	if hasCachedInput {
		cachedInput, err := readCachedInput(session.sessionID, year, day)
		if err != nil {
			return in, err
		}
		return cachedInput, nil
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day), nil)
	if err != nil {
		return in, err
	}

	req.Header.Set("cookie", "session=" + session.sessionID)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return in, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return in, err
    }

	if resp.StatusCode != 200 {
		return in, fmt.Errorf("could not retrieve input: %d => %s", resp.StatusCode, string(body))
	}

	err = cacheInput(string(body), session.sessionID, year, day)
	if err != nil {
		return in, err
	}

	in.rawInput = string(body)

	return in, nil
}
