package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		http.Error(w, "Could not get hostname", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "Hello from: ", hostname)
	fmt.Fprintln(w, "Random number you have requested: ", rand.Intn(1000))
}

func main() {
	fmt.Println("Starting server on port 8080")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
