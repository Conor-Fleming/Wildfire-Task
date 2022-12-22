package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
		first, last := getNameData()
		w.Write([]byte(getJokeData(first, last)))
	})

	// Create the http server at localhost:8080
	server := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	//Start the server
	server.ListenAndServe()
}

func getNameData() (string, string) {
	//Calling Api for the random name
	nameResp, err := http.Get("https://names.mcquay.me/api/v0/")
	if err != nil {
		fmt.Println("Getting name api response: ", err)
	}
	defer nameResp.Body.Close()

	//reading response to byte slice
	nameRespData, err := ioutil.ReadAll(nameResp.Body)
	if err != nil {
		log.Fatal("Read Response Data: ", err)
	}

	// Unmarshalling JSON to name variable of type NameResponse (declared above)
	var name NameResponse
	err = json.Unmarshal(nameRespData, &name)
	if err != nil {
		log.Print("Unmarshal Name: ", err)
	}

	return name.FirstName, name.LastName
}

func getJokeData(first, last string) string {
	//formatting url with given first and last name
	jokeURL := fmt.Sprintf("http://joke.loc8u.com:8888/joke?limitTo=nerdy&firstName=%s&lastName=%s", first, last)

	//Calling Api for the random joke
	jokeResp, err := http.Get(jokeURL)
	if err != nil {
		log.Print("Getting joke api response: ", err)
	}
	defer jokeResp.Body.Close()

	//reading response to byte slice
	jokeRespData, err := ioutil.ReadAll(jokeResp.Body)
	if err != nil {
		log.Print("Read Response Data: ", err)
	}

	//Unmarshalling JSON to joke variable of type JokeResponse
	var joke JokeResponse
	err = json.Unmarshal(jokeRespData, &joke)
	if err != nil {
		log.Print("Unmarshal Joke: ", err)
	}

	output := joke.Value.Joke + "\n"

	return output
}
