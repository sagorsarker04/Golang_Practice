package main

import (
	"log"
	"net/http"
	repository "service/Repository"
	handler "service/handlers"
	user_service "service/service"

	"github.com/gorilla/mux"
)

func main() {
	data := repository.NewGetData()
	svc := user_service.NewUserService(data)
	h := handler.NewUserHandler(svc)
	r := mux.NewRouter()
	r.HandleFunc("/getuser/{id}", h.GetUser)
	log.Fatal(http.ListenAndServe(":8080", r))
}
