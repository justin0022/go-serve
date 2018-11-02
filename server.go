package main

import (
	"log"
	"net/http"
)

func githubHandler() {

}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/github/", githubHandler)

	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
