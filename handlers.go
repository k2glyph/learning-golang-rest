package main

import (
	"encoding/json"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("Hello, there's no one at home!\n"))
	return nil
}

func PaymentFromPaypal(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("Payment Process from paypal begin"))
	// PayToPaypal()
	return nil
}
func PaymentFromStripe(w http.ResponseWriter, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	var request RequestStripe
	err := decoder.Decode(&request)
	if err != nil {
		w.Write([]byte(err.Error()))
		return nil
	}

	// HTTP HEADER SETUP
	w.Header().Add("Content-Type", "application/json;charset=UTF-8")

	// http.StatusOk = 200
	response, err := PaymentStripe(&request)
	if err != nil {
		w.Write([]byte(err.Error()))
		return nil
	}
	// create JSON
	chargeJSON, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte(err.Error()))
		return nil
	}
	// send back JSON
	w.Write(chargeJSON)
	w.WriteHeader(200)
	return nil
}

func responseWithError(err error) []byte {
	if err != nil {
		// our custom error type
		return []byte(err.Error())
	}
	return nil
}
