package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// set env vars
	env()
	// register a NewRouter() from router.go gorilla mux lib
	router := NewRouter()
	SetStripeKey(os.Getenv("STRIPE_PRIVATE_KEY"))
	fmt.Println("Server is running on " + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))

}
