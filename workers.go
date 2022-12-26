package main

import (
	"errors"
	"strings"
)

func startWorkers() (string, error) {

	// TODO:
	// - Look into error handling with channels.
	// - Should these channels be buffered?

	nameChan := make(chan *NameResponse)
	jokeChan := make(chan *JokeResponse)

	go NameWorker(nameChan)
	go JokeWorker(jokeChan)

	name := <-nameChan
	joke := <-jokeChan

	//doing it this way so the JokeData() func doesnt have to wait for the NameData values before we call it.
	if name != nil && joke != nil {
		// Replace name values in Joke with values from Name API
		joke.Value.Joke = strings.ReplaceAll(joke.Value.Joke, "*first", name.FirstName)
		joke.Value.Joke = strings.ReplaceAll(joke.Value.Joke, "*last", name.LastName)

		return joke.Value.Joke, nil
	}
	return "", errors.New("missing data")
}
