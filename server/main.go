package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("msg")
	if msg == "" {
		msg = "Nothing received (did you forget ?msg=... ?)"
	}
	fmt.Println("Received:", msg)
	fmt.Fprintf(w, "Hello Client, I got: %s", msg)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server running at :8080")
	http.ListenAndServe(":8080", nil)

}
