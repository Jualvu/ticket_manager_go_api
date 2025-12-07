package handlers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/jualvu/go-tickets-api/internal/store"
	"github.com/jualvu/go-tickets-api/internal/models"
	"github.com/jualvu/go-tickets-api/internal/models/dto/ticket_dto"
	"strconv"
)

type TicketHandler struct {
	ticketStore store.TicketStore
}

func NewTicketHandler (ticketStore store.TicketStore) *TicketHandler {
	return &TicketHandler{ticketStore: ticketStore} // create and return TicketHandler pointer with assigned ticketStore 
}	

func (h *TicketHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return 
	}

	w.Header().Set("Content-Type", "application/json") // set the response content

	res := map[string]any{
		"success": true,
		"ticketID": -1,
		"message": "",
	}

	var createTicketRequest ticket_dto.CreateTicketRequest
	err := json.NewDecoder(r.Body).Decode(&createTicketRequest); if err != nil {
		http.Error(w, "Invalid body content", http.StatusBadRequest)
		return
	}

	newTicketID, err := h.ticketStore.Create(createTicketRequest)
	if err != nil {
		http.Error(w, "Error when trying to create ticket.", http.StatusInternalServerError)
		return
	}

	res["message"] = "Successfully created ticket."
	res["ticketID"] = newTicketID

	json.NewEncoder(w).Encode(res)
	return
} 

func (h *TicketHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return 
	}

	w.Header().Set("Content-Type", "application/json") // set the response content

	res := map[string]any{
		"success": true,
		"message": "",
	}

	var updateTicketRequest ticket_dto.UpdateTicketRequest
	err := json.NewDecoder(r.Body).Decode(&updateTicketRequest); if err != nil {
		http.Error(w, "Invalid body content", http.StatusBadRequest)
		return
	}

	err = h.ticketStore.Update(updateTicketRequest)
	if err != nil {
		msg := fmt.Sprintf("Error when trying to update ticket with id [%v].", updateTicketRequest.Id)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	res["message"] = fmt.Sprintf("Successfully updated ticket [%v].", updateTicketRequest.Id)

	json.NewEncoder(w).Encode(res)
	return
} 

func (h *TicketHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return 
	}

	w.Header().Set("Content-Type", "application/json") // set the response content

	res := map[string]any{
		"success": true,
		"message": "",
	}

	var deleteTicketRequest ticket_dto.DeleteTicketRequest
	err := json.NewDecoder(r.Body).Decode(&deleteTicketRequest); if err != nil {
		http.Error(w, "Invalid body content", http.StatusBadRequest)
		return
	}

	err = h.ticketStore.Delete(deleteTicketRequest)
	if err != nil {
		msg := fmt.Sprintf("Error when trying to delete ticket with id [%v].", deleteTicketRequest.Id)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	res["message"] = fmt.Sprintf("Successfully deleted ticket [%v].", deleteTicketRequest.Id)

	json.NewEncoder(w).Encode(res)
	return
} 

func (h *TicketHandler) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return 
	}

	w.Header().Set("Content-Type", "application/json") // set the response content

	var tickets []*models.Ticket
	
	res := map[string]any{
		"success": true,
		"tickets": tickets,
		"message": "",
	}

	id_string := r.URL.Query().Get("id") // Check if a specific ticket is being requested
	
	if id_string != "" {
		id, err := strconv.Atoi(id_string)

		if err != nil {
			res["success"] = false
			res["message"] = fmt.Sprintf("Error when trying to parse ID provided: [%v].", id_string)
			json.NewEncoder(w).Encode(res)
			return
		}

		ticketFound, err := h.ticketStore.GetByID(id)
		if err != nil {
			res["success"] = false
			res["message"] = fmt.Sprintf("Ticket with ID [%v] not found.", id)
			json.NewEncoder(w).Encode(res)
			return
		}

		tickets = append(tickets, ticketFound)
		res["message"] = fmt.Sprintf("Success: Ticket with ID [%v] found.", id)

		} else {

		all_tickets, err := h.ticketStore.GetAll()
		if err != nil {
			http.Error(w, "Error when trying to get all tickets.", http.StatusInternalServerError)
			return
		}

		tickets = all_tickets
		res["message"] = "Success: All tickets retrieved."
	}

	res["tickets"] = tickets
	json.NewEncoder(w).Encode(res)
	return
} 
