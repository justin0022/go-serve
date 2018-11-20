package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getUser(username string) string {
	response, _ := http.Get("https://api.github.com/users/" + username)
	responseData, _ := ioutil.ReadAll(response.Body)
	userData := string(responseData)
	return userData
}

func GithubHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	username := params["username"]
	w.WriteHeader(http.StatusOK)
	userData := getUser(username)
	log.Println(userData)
}

func main() {
	router := mux.NewRouter()
	fs := http.FileServer(http.Dir("static"))
	router.Handle("/", fs)
	router.HandleFunc("/github/{username}", GithubHandler)
	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", router)
}
