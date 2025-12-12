package handlers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/jualvu/go-tickets-api/internal/store"
	"github.com/jualvu/go-tickets-api/internal/models/dto/auth_dto"
)

type AuthHandler struct {
	userStore store.UserStore
}

func NewAuthHandler (userStore store.UserStore) *AuthHandler {
	return &AuthHandler{ userStore: userStore }
}

func (h *AuthHandler) Login (w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return 
	}

	w.Header().Set("Content-Type", "application/json")

	res := map[string]any{
		"success": true,
		"message": "",
		"user": nil,
	}

	var loginAuthRequest auth_dto.LoginAuthRequest
	err := json.NewDecoder(r.Body).Decode(&loginAuthRequest); if err != nil {
		http.Error(w, "Invalid body content", http.StatusBadRequest)
		return
	}

	userFound, err := h.userStore.GetByEmailAndPassword(loginAuthRequest.Email, loginAuthRequest.Password)
	if err != nil {
		res["success"] = false
		res["message"] = fmt.Sprintf("User with Email [%v] and Password[%v] not found.", loginAuthRequest.Email, loginAuthRequest.Password)
		json.NewEncoder(w).Encode(res)
		return
	}

	res["message"] = fmt.Sprintf("Success: User with Email [%v] and Password [%v] found.", loginAuthRequest.Email, loginAuthRequest.Password)
	res["user"] = userFound

	json.NewEncoder(w).Encode(res)
	return
}

