package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Home Page")
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "About Page")
	})

	http.HandleFunc("/method-check", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			fmt.Fprintln(w, "This is POST method")
		} else {
			fmt.Fprintln(w, "This is get method")
		}
	})

	fmt.Println("Server is running on http://localhost:8081")
	http.ListenAndServe(":8081", nil)
}
