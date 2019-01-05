package main

import "net/http"

func hello(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("Hello, there's no one at home!\n"))
	return nil
}
