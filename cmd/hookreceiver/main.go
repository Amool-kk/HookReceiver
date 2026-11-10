package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("Starting the server")

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Hook Receiver"))
	})

	server := http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: router,
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("Failed to start server")
	}
}
