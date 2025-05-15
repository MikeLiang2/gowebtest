package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("msg")
	fmt.Println("Received:", msg)
	fmt.Fprintf(w, "Server received your message: %s", msg)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
