package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var POD_NAME = os.Getenv("POD_NAME")

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("I'm %s", POD_NAME)))
}

func main() {
	http.HandleFunc("/", handler)

	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
		IdleTimeout:  5 * time.Second,
	}

	log.Println("Server listening on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
