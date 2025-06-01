package handler

import (
	"encoding/json"
	"net/http"
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
		return
	}
	json.NewEncoder(w).Encode(user)
}
