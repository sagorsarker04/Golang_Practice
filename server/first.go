package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Handle root URL
	//w http.ResponseWriter is used for sending the response to the client
	//r is used for bringing the response from the client
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
