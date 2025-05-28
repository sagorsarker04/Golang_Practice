package main

import (
	"log"
	"net/http"
	"test/service"
	"test/handlers"
	"github.com/gorilla/mux"
)

func main() {
	svc := service.NewUserService()  // service instance
	h := handler.NewUserHandler(svc) // handler with injected service

	r := mux.NewRouter()
	r.HandleFunc("/users/{id}", h.GetUserHandler).Methods("GET")

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
