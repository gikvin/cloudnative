package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		http.Error(w, "Could not get hostname", http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, "Hello from: ", hostname)
}

func main() {
	fmt.Println("Starting server on port 8080")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
