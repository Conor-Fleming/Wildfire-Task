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
}

func getNameData() (string, string) {
	//Calling Api for the random name
	nameResp, err := http.Get("https://names.mcquay.me/api/v0/")
	if err != nil {
		log.Fatal("Getting name api response", err)
	}
	defer nameResp.Body.Close()

	nameRespData, err := ioutil.ReadAll(nameResp.Body)
	if err != nil {
		log.Fatal("Read Response Data", err)
	}

	var name NameResponse
	err = json.Unmarshal(nameRespData, &name)
	if err != nil {
		log.Fatal("Unmarshal Name", err)
	}

	return name.FirstName, name.LastName
}

func getJokeData(first, last string) {
	//formatting url with given first and last name
	jokeURL := fmt.Sprintf("http://joke.loc8u.com:8888/joke?limitTo=nerdy&firstName=%s&lastName=%s", first, last)

	//Calling Api for the random name
	jokeResp, err := http.Get(jokeURL)
	if err != nil {
		log.Fatal("Getting name api response", err)
	}
	defer jokeResp.Body.Close()

	jokeRespData, err := ioutil.ReadAll(jokeResp.Body)
	if err != nil {
		log.Fatal("Read Response Data", err)
	}

	var joke JokeResponse
	err = json.Unmarshal(jokeRespData, &joke)
	if err != nil {
		log.Fatal("Unmarshal Name", err)
	}

	fmt.Println(joke.Value.Joke)
}
