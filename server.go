package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func main() {
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "Dinesh", Lastname: "Katwal", Address: &Address{City: "Kathmandu", State: "Bagmati"}})
	people = append(people, Person{ID: "2", Firstname: "Dilip", Lastname: "Katwal"})
	people = append(people, Person{ID: "3", Firstname: "Mina", Lastname: "Ma", Address: &Address{City: "Lalbandi", State: "Janakpur"}})
	people = append(people, Person{ID: "4", Firstname: "Shisir", Address: &Address{City: "Beni", State: "Gandaki"}})

	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	log.Fatal(http.ListenAndServe(":8082", router))
}
