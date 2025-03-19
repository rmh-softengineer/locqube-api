package main

import (
	"net/http"
	"time"
)

func main() {
	timeoutValue := 60

	timeout := time.Duration(timeoutValue) * time.Second

	server := http.Server{
		Addr:         ":8080",
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}

	server.ListenAndServe()
}
