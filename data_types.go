package main

type NameResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type JokeResponse struct {
	Value struct {
		Joke string `json:"joke"`
	} `json:"value"`
}
