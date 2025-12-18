package handlers

import (
	"fmt"
	"strings"
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

	res := map[string]any{
		"success": true,
		"userID": -1,
		"message": "",
	}

	var createUserRequest user_dto.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&createUserRequest); if err != nil {
		http.Error(w, "Invalid body content", http.StatusBadRequest)
		return
	}

	newUserID, err := h.userStore.Create(createUserRequest)
	if err != nil {
		http.Error(w, "Error when trying to create User.", http.StatusInternalServerError)
		return
	}

	res["message"] = "Successfully created user."
	res["userID"] = newUserID

	json.NewEncoder(w).Encode(res)
	return
} 

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return 
	}

	w.Header().Set("Content-Type", "application/json") // set the response content

	res := map[string]any{
		"success": true,
		"message": "",
	}

	id_string := strings.TrimPrefix(r.URL.Path, "/users/")
	if id_string == "" {
		res["success"] = false
		res["message"] = "No user ID provided to Update."
		json.NewEncoder(w).Encode(res)
		return
	}

	id, err := strconv.Atoi(id_string)
	if err != nil {
		res["success"] = false
		res["message"] = fmt.Sprintf("Error when trying to parse ID provided: [%v].", id_string)
		json.NewEncoder(w).Encode(res)
		return
	}

	var updateUserRequest user_dto.UpdateUserRequest
	err = json.NewDecoder(r.Body).Decode(&updateUserRequest); if err != nil {
		http.Error(w, "Invalid body content", http.StatusBadRequest)
		return
	}

	updateUserRequest.Id = id

	err = h.userStore.Update(updateUserRequest)
	if err != nil {
		msg := fmt.Sprintf("Error when trying to update User with id [%v].", updateUserRequest.Id)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	res["message"] = fmt.Sprintf("Successfully updated user [%v].", updateUserRequest.Id)

	json.NewEncoder(w).Encode(res)
	return
} 

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return 
	}

	w.Header().Set("Content-Type", "application/json") // set the response content

	res := map[string]any{
		"success": true,
		"message": "",
	}

	id_string := strings.TrimPrefix(r.URL.Path, "/users/")
	if id_string == "" {
		res["success"] = false
		res["message"] = "No user ID provided to Delete."
		json.NewEncoder(w).Encode(res)
		return
	}

	id, err := strconv.Atoi(id_string)
	if err != nil {
		res["success"] = false
		res["message"] = fmt.Sprintf("Error when trying to parse ID provided: [%v].", id_string)
		json.NewEncoder(w).Encode(res)
		return
	}

	var deleteUserRequest user_dto.DeleteUserRequest
	deleteUserRequest.Id = id

	err = h.userStore.Delete(deleteUserRequest)
	if err != nil {
		msg := fmt.Sprintf("Error when trying to delete User with id [%v].", deleteUserRequest.Id)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	res["message"] = fmt.Sprintf("Successfully deleted user [%v].", deleteUserRequest.Id)

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
	
	id_string := strings.TrimPrefix(r.URL.Path, "/users/")
	
	if id_string == "" {
		all_users, err := h.userStore.GetAll()
		if err != nil {
			http.Error(w, "Error when trying to get all users.", http.StatusInternalServerError)
			return
		}
		users = all_users
		res["message"] = "Success: All user retrieved."

	} else {

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
	}

	res["users"] = users
	json.NewEncoder(w).Encode(res)
	return
} 


