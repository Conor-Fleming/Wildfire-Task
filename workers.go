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

	//unsure about error handling in this case
	//printed message to log and return nil upon error in these two functions so i suppose the error is handled if code falls through on line 47
	go func() {
		name, _ = NameData()
		wg.Done()
	}()
	go func() {
		joke, _ = JokeData()
		wg.Done()
	}()
	wg.Wait()

	if name != nil && joke != nil {
		// Replace name values in Joke with values from Name API
		joke.Value.Joke = strings.ReplaceAll(joke.Value.Joke, "*first", name.FirstName)
		joke.Value.Joke = strings.ReplaceAll(joke.Value.Joke, "*last", name.LastName)

		return joke.Value.Joke, nil
	}
	return "", errors.New("missing data")
}
