package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PATCH")
	r.HandleFunc("/users/{id}", DeleteUser).Methods("delete")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	initializeRouter()
}
