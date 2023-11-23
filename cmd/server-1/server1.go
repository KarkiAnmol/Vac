package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "pong")
	})

	fmt.Println("Server 1 listening on :8081...")
	http.ListenAndServe(":8081", nil)
}
