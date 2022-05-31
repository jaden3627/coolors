package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./views"))
	http.Handle("/", fs)

	log.Print("Listening on :8010...")
	err := http.ListenAndServe(":8010", nil)
	if err != nil {
		log.Fatal(err)
	}
}
