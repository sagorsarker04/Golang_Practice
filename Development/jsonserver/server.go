package jsonserver

import (
	"fmt"
	"encoding/json"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func SendJSON() {
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		user := User{Name: "Sagor", Email: "sagor@example.com"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	})

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
