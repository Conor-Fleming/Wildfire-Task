package main

import (
	"log"
	"net/http"
)

func main() {
	// Set up handler function for the "/" route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res, err := startWorkers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write([]byte(res))
	})

	// Create the http server at localhost:8080
	server := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	//Start the server, if port fails to open log problem and exit
	log.Fatal(server.ListenAndServe())
}
