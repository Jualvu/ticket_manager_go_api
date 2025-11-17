package handlers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/jualvu/go-tickets-api/internal/store"
	"github.com/jualvu/go-tickets-api/internal/models/dto/user_dto"
)

type UserHandler struct {
	userStore store.UserStore
}

func NewUserHandler (userStore store.UserStore) *UserHandler {
	return &UserHandler{userStore: userStore} // create and return UserHandler pointer with assigned userStore 
}	

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return 
	}

	w.Header().Set("Content-Type", "application/json") // set the response content

	var createUserRequest user_dto.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&createUserRequest); if err != nil {
		http.Error(w, "Invalid body content", http.StatusBadRequest)
		return
	}

	fmt.Printf(createUserRequest.Name) // test

	newUserID, err := h.userStore.Create(createUserRequest)
	if err != nil {
		http.Error(w, "Error when trying to create User.", http.StatusInternalServerError)
		return
	}

	res := map[string]any{
		"success": true,
		"user_id": newUserID,
	}

	json.NewEncoder(w).Encode(res)
	return
} 


func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return 
	}

	w.Header().Set("Content-Type", "application/json") // set the response content

	users, err := h.userStore.GetAll()
	if err != nil {
		http.Error(w, "Error when trying to get all users.", http.StatusInternalServerError)
		return
	}

	res := map[string]any{
		"success": true,
		"user_id": users,
	}

	json.NewEncoder(w).Encode(res)
	return
} 