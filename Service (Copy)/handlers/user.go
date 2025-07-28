package handler

import (
	"encoding/json"
	"net/http"
	"service/model"
	user_service "service/service"
	"strconv"

	"github.com/gorilla/mux"
)

type UserHanlder struct {
	service user_service.UserService
}

func NewUserHandler(s user_service.UserService) *UserHanlder {
	return &UserHanlder{service: s}
}
func (h *UserHanlder) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	idstr, err := strconv.Atoi(idStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user, err := h.service.GetUser(idstr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Failed to get User", http.StatusBadGateway)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (h *UserHanlder) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Failed to Decode json2", http.StatusBadGateway)
		return
	}

	userResponse, err := h.service.CreateUser(user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Failed to get response form the sevice", http.StatusBadGateway)
		return
	}
	json.NewEncoder(w).Encode(userResponse)

}

func (h *UserHanlder) CreateTable(w http.ResponseWriter, r *http.Request) {
	err := h.service.CreateTable()
	if err != nil {
		http.Error(w, "Service layer a somossa", http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode("Table created Successfully")
}

func (h *UserHanlder) DropTable(w http.ResponseWriter, r *http.Request) {
	err := h.service.DropTable()
	if err != nil {
		http.Error(w, "Failed to drop table", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Table dropped successfully")
}
