package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GithubHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	w.WriteHeader(http.StatusOK)
	response, _ := http.Get("https://api.github.com/users/" + params["username"])
	responseData, _ := ioutil.ReadAll(response.Body)
	log.Println(string(responseData))
}

func main() {
	router := mux.NewRouter()
	fs := http.FileServer(http.Dir("static"))
	router.Handle("/", fs)
	router.HandleFunc("/github/{username}", GithubHandler)
	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", router)
}
