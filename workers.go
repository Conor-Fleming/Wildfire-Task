package main

import (
	"errors"
	"strings"
	"sync"
)

func startWorkers() (string, error) {
	var wg sync.WaitGroup
	wg.Add(2)

	var name *NameResponse
	var joke *JokeResponse

	// Errors are not used here as it seemed redundant.
	// If there is an error returned from either NameData() or JokeData()
	// the code will fall through on line 30 and return the appropriate error
	go func() {
		name, _ = NameData()
		wg.Done()
	}()
	go func() {
		joke, _ = JokeData()
		wg.Done()
	}()
	wg.Wait()

	//doing it this way so the JokeData() func doesnt have to wait for the NameData values before we call it.
	if name != nil && joke != nil {
		// Replace name values in Joke with values from Name API
		joke.Value.Joke = strings.ReplaceAll(joke.Value.Joke, "*first", name.FirstName)
		joke.Value.Joke = strings.ReplaceAll(joke.Value.Joke, "*last", name.LastName)

		return joke.Value.Joke, nil
	}
	return "", errors.New("missing data")
}
