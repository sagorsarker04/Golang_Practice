package main

import (
	"log"
	"net/http"
	repository "service/Repository"
	"service/config"
	database "service/database"
	handler "service/handlers"
	user_service "service/service"

	"github.com/gorilla/mux"
)

func main() {

	cfg := config.LoadConfig()
	db := database.InitDB(cfg)

	defer db.Close()
	data := repository.NewGetData(db)
	svc := user_service.NewUserService(data)
	h := handler.NewUserHandler(svc)

	r := mux.NewRouter()
	r.HandleFunc("/users/{id}", h.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/users", h.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/tables/users", h.CreateTable).Methods(http.MethodPost)
	r.HandleFunc("/tables/users", h.DropTable).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":8080", r))
}
