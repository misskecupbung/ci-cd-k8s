package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	r := mux.NewRouter()
	r.HandleFunc("/api/healthcheck", healthy)

	http.Handle("/api/healthcheck", r)
	fmt.Println("Starting up on " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func healthy(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "I'm healthy!")
}
