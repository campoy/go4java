package main

import (
	"fmt"
	"net/http"
)

var nextID = 0

func handler(w http.ResponseWriter, q *http.Request) {
	fmt.Fprintf(w, "<h1>You got %v<h1>", nextID)
	nextID++
}

func main() {
	http.HandleFunc("/next", handler)
	http.ListenAndServe("localhost:8080", nil)
}
