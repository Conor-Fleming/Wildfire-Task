package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func NameWorker(c chan<- *NameResponse) {
	//Calling Api for the random name
	nameResp, err := http.Get("https://names.mcquay.me/api/v0/")
	if err != nil {
		log.Print("Getting name api response: ", err)
	}
	defer nameResp.Body.Close()

	//reading response to byte slice
	nameRespData, err := ioutil.ReadAll(nameResp.Body)
	if err != nil {
		log.Print("Read Response Data: ", err)
	}

	// Unmarshalling JSON to name variable of type NameResponse (declared above)
	name := &NameResponse{}
	err = json.Unmarshal(nameRespData, &name)
	if err != nil {
		log.Print("Unmarshal Name: ", err)
	}

	c <- name
}

func JokeWorker(c chan<- *JokeResponse) {
	// TODO: use url package to build these strings properly

	//formatting url with given first and last name
	jokeURL := fmt.Sprintf("http://joke.loc8u.com:8888/joke?limitTo=nerdy&firstName=*first&lastName=*last")

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
	joke := &JokeResponse{}
	err = json.Unmarshal(jokeRespData, &joke)
	if err != nil {
		log.Print("Unmarshal Joke: ", err)
	}

	c <- joke
}
