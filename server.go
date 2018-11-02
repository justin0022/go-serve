package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func githubHandler(w http.ResponseWriter, req *http.Request) {

}

func main() {
	router := mux.NewRouter()
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	router.HandleFunc("/github/", githubHandler).Methods("GET")

	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
