package urlparameter

import (
	"fmt"
	"net/http"
)

func Parameter() {
	http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Path[len("/user/"):]
		fmt.Fprintf(w, "Hello, %s!", username)
	})

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
