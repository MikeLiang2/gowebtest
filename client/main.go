package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// This is a simple HTTP client that sends a GET request to the server
func main() {
	resp, err := http.Get("http://192.168.1.100:8080/?msg=PingFromClient")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Server replied:", string(body))
}
