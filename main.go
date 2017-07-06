package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	h := http.NewServeMux()
	h.HandleFunc("/", index)
	err := http.ListenAndServe(":"+port, h)
	if err != nil {
		log.Fatal(err)
	}
}
