package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/users", CreatetUsers).Methods("POST")
	router.HandleFunc("/users/{id}", UpdateUsers).Methods("PATCH")
	router.HandleFunc("/users/{id}", DeleteUsers).Methods("delete")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	initializeRouter()
}
