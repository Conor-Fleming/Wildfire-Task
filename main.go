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
		Id   int    `json:"id"`
		Joke string `json:"joke"`
	} `json:"value"`
}

func main() {
	first, last := getNameData()
	getJokeData(first, last)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(getJokeData(first, last)))
	})

	http.ListenAndServe(":8080", nil)
}

func getNameData() (string, string) {
	//Calling Api for the random name
	nameResp, err := http.Get("https://names.mcquay.me/api/v0/")
	if err != nil {
		log.Fatal("Getting name api response", err)
	}
	defer nameResp.Body.Close()

	//reading response to byte slice
	nameRespData, err := ioutil.ReadAll(nameResp.Body)
	if err != nil {
		log.Fatal("Read Response Data", err)
	}

	// Unmarshalling JSON to name variable of type NameResponse (declared above)
	var name NameResponse
	err = json.Unmarshal(nameRespData, &name)
	if err != nil {
		log.Fatal("Unmarshal Name", err)
	}

	return name.FirstName, name.LastName
}

func getJokeData(first, last string) string {
	//formatting url with given first and last name
	jokeURL := fmt.Sprintf("http://joke.loc8u.com:8888/joke?limitTo=nerdy&firstName=%s&lastName=%s", first, last)

	//Calling Api for the random name
	jokeResp, err := http.Get(jokeURL)
	if err != nil {
		log.Fatal("Getting name api response", err)
	}
	defer jokeResp.Body.Close()

	//reading response to byte slice
	jokeRespData, err := ioutil.ReadAll(jokeResp.Body)
	if err != nil {
		log.Fatal("Read Response Data", err)
	}

	//Unmarshalling JSON to joke variable of type JokeResponse
	var joke JokeResponse
	err = json.Unmarshal(jokeRespData, &joke)
	if err != nil {
		log.Fatal("Unmarshal Name", err)
	}

	return joke.Value.Joke
}
