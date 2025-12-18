package handlers

import (
	"fmt"
	"strings"
	"net/http"
	"encoding/json"
	"github.com/jualvu/go-tickets-api/internal/store"
	"github.com/jualvu/go-tickets-api/internal/models"
	"github.com/jualvu/go-tickets-api/internal/models/dto/comment_dto"
	"strconv"
)

type CommentHandler struct {
	commentStore store.CommentStore
}

func NewCommentHandler (commentStore store.CommentStore) *CommentHandler {
	return &CommentHandler{commentStore: commentStore} // create and return CommentHandler pointer with assigned commentStore
}

func (h *CommentHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json") // set the response content

	res := map[string]any{
		"success": true,
		"commentID": -1,
		"message": "",
	}

	var createCommentRequest comment_dto.CreateCommentRequest
	err := json.NewDecoder(r.Body).Decode(&createCommentRequest); if err != nil {
		http.Error(w, "Invalid body content", http.StatusBadRequest)
		return
	}

	newCommentID, err := h.commentStore.Create(createCommentRequest)
	if err != nil {
		http.Error(w, "Error when trying to create Comment.", http.StatusInternalServerError)
		return
	}

	res["message"] = "Successfully created comment."
	res["commentID"] = newCommentID

	json.NewEncoder(w).Encode(res)
	return
}

func (h *CommentHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json") // set the response content

	res := map[string]any{
		"success": true,
		"message": "",
	}

	id_string := strings.TrimPrefix(r.URL.Path, "/comments/")
	if id_string == "" {
		res["success"] = false
		res["message"] = "No comment ID provided to Update."
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

	var updateCommentRequest comment_dto.UpdateCommentRequest
	err = json.NewDecoder(r.Body).Decode(&updateCommentRequest); if err != nil {
		http.Error(w, "Invalid body content", http.StatusBadRequest)
		return
	}
	updateCommentRequest.Id = id

	err = h.commentStore.Update(updateCommentRequest)
	if err != nil {
		msg := fmt.Sprintf("Error when trying to update comment with id [%v].", updateCommentRequest.Id)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	res["message"] = fmt.Sprintf("Successfully updated comment with id [%v].", updateCommentRequest.Id)

	json.NewEncoder(w).Encode(res)
	return
}

func (h *CommentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json") // set the response content

	res := map[string]any{
		"success": true,
		"message": "",
	}

	id_string := strings.TrimPrefix(r.URL.Path, "/comments/")
	if id_string == "" {
		res["success"] = false
		res["message"] = "No comment ID provided to Delete."
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

	var deleteCommentRequest comment_dto.DeleteCommentRequest
	deleteCommentRequest.Id = id 

	err = h.commentStore.Delete(deleteCommentRequest)
	if err != nil {
		msg := fmt.Sprintf("Error when trying to delete comment with id [%v].", deleteCommentRequest.Id)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}	

	res["message"] = fmt.Sprintf("Successfully deleted comment with id [%v].", deleteCommentRequest.Id)

	json.NewEncoder(w).Encode(res)
	return
}	

func (h *CommentHandler) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json") // set the response content

	var comments []*models.Comment

	res := map[string]any{
		"success": true,
		"comments": comments,
		"message": "",
	}

	id_string := strings.TrimPrefix(r.URL.Path, "/comments/")

	if id_string == "" {
		all_comments, err := h.commentStore.GetAll()
		if err != nil {
			res["success"] = false
			http.Error(w, "Error when trying to get all comments.", http.StatusInternalServerError)
			return
		}	
		comments = all_comments
		res["message"] = "Success: All comments retrieved."

	} else {

		id, err := strconv.Atoi(id_string)
		if err != nil {
			res["success"] = false
			res["message"] = fmt.Sprintf("Error when trying to parse ID provided: [%v].", id_string)
			json.NewEncoder(w).Encode(res)
			return
		}

		commentFound, err := h.commentStore.GetByID(id)
		if err != nil {
			res["success"] = false
			res["message"] = fmt.Sprintf("Comment with ID [%v] not found.", id)
			json.NewEncoder(w).Encode(res)
			return
		}

		comments = append(comments, commentFound)
		res["comments"] = fmt.Sprintf("Success: Comment with ID [%v] found.", id)
	}
 	
	res["comments"] = comments
	json.NewEncoder(w).Encode(res)

	return
}	