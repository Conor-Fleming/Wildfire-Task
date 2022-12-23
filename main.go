package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

type NameResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type JokeResponse struct {
	Value struct {
		Joke string `json:"joke"`
	} `json:"value"`
}

func main() {
	// Set up handler function for the "/" route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Get first and last names from getNameData() and use those as args in getJokeData()
		// write the returned joke to the response
		var wg sync.WaitGroup
		wg.Add(2)

		var name *NameResponse
		var joke *JokeResponse

		//unsure about error handling in this case
		//printed message to log and return nil upon error in these two functions so i suppose the error is handled if code falls through on line 47
		go func() {
			name, _ = getNameData()
			wg.Done()
		}()
		go func() {
			joke, _ = getJokeData()
			wg.Done()
		}()
		wg.Wait()

		if name != nil && joke != nil {
			joke.Value.Joke = strings.Replace(joke.Value.Joke, "--first--", name.FirstName, 1)
			joke.Value.Joke = strings.Replace(joke.Value.Joke, "--last--", name.LastName, 1)

			w.Write([]byte(joke.Value.Joke))
		}
		http.Error(w, "There was an error getting the data", http.StatusInternalServerError)
	})

	// Create the http server at localhost:8080
	server := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	//Start the server
	server.ListenAndServe()
}

func getNameData() (*NameResponse, error) {
	//Calling Api for the random name
	nameResp, err := http.Get("https://names.mcquay.me/api/v0/")
	if err != nil {
		log.Print("Getting name api response: ", err)
		return nil, err
	}
	defer nameResp.Body.Close()

	//reading response to byte slice
	nameRespData, err := ioutil.ReadAll(nameResp.Body)
	if err != nil {
		log.Print("Read Response Data: ", err)
		return nil, err
	}

	// Unmarshalling JSON to name variable of type NameResponse (declared above)
	name := &NameResponse{}
	err = json.Unmarshal(nameRespData, &name)
	if err != nil {
		log.Print("Unmarshal Name: ", err)
		return nil, err
	}

	return name, nil
}

func getJokeData() (*JokeResponse, error) {
	//formatting url with given first and last name
	jokeURL := fmt.Sprintf("http://joke.loc8u.com:8888/joke?limitTo=nerdy&firstName=--first--&lastName=--last--")

	//Calling Api for the random joke
	jokeResp, err := http.Get(jokeURL)
	if err != nil {
		log.Print("Getting joke api response: ", err)
		return nil, err
	}
	defer jokeResp.Body.Close()

	//reading response to byte slice
	jokeRespData, err := ioutil.ReadAll(jokeResp.Body)
	if err != nil {
		log.Print("Read Response Data: ", err)
		return nil, err
	}

	//Unmarshalling JSON to joke variable of type JokeResponse
	joke := &JokeResponse{}
	err = json.Unmarshal(jokeRespData, &joke)
	if err != nil {
		log.Print("Unmarshal Joke: ", err)
		return nil, err
	}

	return joke, nil
}
