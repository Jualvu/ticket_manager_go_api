package handlers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/jualvu/go-tickets-api/internal/store"
	"github.com/jualvu/go-tickets-api/internal/models"
	"github.com/jualvu/go-tickets-api/internal/models/dto/user_dto"
	"strconv"
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

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return 
	}

	w.Header().Set("Content-Type", "application/json") // set the response content

	var users []*models.User
	
	res := map[string]any{
		"success": true,
		"users": users,
		"message": "",
	}
	
	id_string := r.URL.Query().Get("id") // Check if a specific user is being requested
	
	if id_string != "" {
		id, err := strconv.Atoi(id_string)

		if err != nil {
			res["success"] = false
			res["message"] = fmt.Sprintf("Error when trying to parse ID provided: [%v].", id_string)
			json.NewEncoder(w).Encode(res)
			return
		}

		userFound, err := h.userStore.GetByID(id)
		if err != nil {
			res["success"] = false
			res["message"] = fmt.Sprintf("User with ID [%v] not found.", id)
			json.NewEncoder(w).Encode(res)
			return
		}

		users = append(users, userFound)
		res["message"] = fmt.Sprintf("Success: User with ID [%v] found.", id)

	} else {
		all_users, err := h.userStore.GetAll()
		if err != nil {
			http.Error(w, "Error when trying to get all users.", http.StatusInternalServerError)
			return
		}
		users = all_users
		res["message"] = "Success: All user retrieved."
	}

	res["users"] = users
	json.NewEncoder(w).Encode(res)
	return
} 


