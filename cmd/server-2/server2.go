package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Send a ping request to Server 1
	resp, err := http.Get("http://localhost:8081/get-json")
	if err != nil {
		fmt.Println("Error sending ping request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Display the pong received from Server 1
	fmt.Println("Received pong from Server 1:", string(body))
}
