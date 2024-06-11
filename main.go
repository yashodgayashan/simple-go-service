package main

import (
	"log"
	"net/http"
	"simple-go/handler"
)

func main() {
	http.HandleFunc("/hello", handler.HelloHandler)
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("could not start server: %s", err.Error())
	}
}
