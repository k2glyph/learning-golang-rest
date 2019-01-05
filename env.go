package main

import "os"

func env() {
	whichEnv := os.Getenv("PAYMENT_PROCESS_ENV")
	os.Setenv("PORT", "8081")
	if whichEnv == "local" || whichEnv == "" {
		os.Setenv("PORT", "8081")
	}
}
