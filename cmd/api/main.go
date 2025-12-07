package main 

import (
	"log"
	"net/http"
	"fmt"
	"database/sql"
	"github.com/jualvu/go-tickets-api/internal/store"
	"github.com/jualvu/go-tickets-api/internal/handlers"
)

func main() {

	// basically define each endpoint here and call the handlers on each method type
	// we need to manage users - tickets - comments

	dbConnection, err := sql.Open("sqlite3", "./internal/database/ticket_system.db")
	if err != nil {
		fmt.Errorf("Error when trying to open the Database ticket_system.db")
	}
	defer dbConnection.Close()

	mux := http.NewServeMux()

	// user store

	userStore := store.NewUserStore(dbConnection)
	userHandler := handlers.NewUserHandler(userStore)
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			userHandler.Get(w, r)
			break
		case http.MethodPost:
			userHandler.Create(w, r)
			break
		case http.MethodPut:
			userHandler.Update(w, r)
			break
		case http.MethodDelete:
			userHandler.Delete(w, r)
			break
		default:
			http.Error(w, "Method not allowed for /users endpoint.", http.StatusMethodNotAllowed)
		}
	})

	// ticket store

	ticketStore := store.NewTicketStore(dbConnection)
	ticketHandler := handlers.NewTicketHandler(ticketStore)
	mux.HandleFunc("/tickets", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			ticketHandler.Get(w, r)
			break
		case http.MethodPost:
			ticketHandler.Create(w, r)
			break
		case http.MethodPut:
			ticketHandler.Update(w, r)
			break
		case http.MethodDelete:
			ticketHandler.Delete(w, r)
			break
		default:
			http.Error(w, "Method not allowed for /tickets endpoint.", http.StatusMethodNotAllowed)
		}
	})

	// comment store

	commentStore := store.NewCommentStore(dbConnection)
	commentHandler := handlers.NewCommentHandler(commentStore)
	mux.HandleFunc("/comments", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			commentHandler.Get(w, r)
			break
		case http.MethodPost:
			commentHandler.Create(w, r)
			break
		case http.MethodPut:
			commentHandler.Update(w, r)
			break
		case http.MethodDelete:
			commentHandler.Delete(w, r)
			break
		default:
			http.Error(w, "Method not allowed for /comments endpoint.", http.StatusMethodNotAllowed)
		}
	})
	

	log.Println("API Server listening on PORT :8080")
	http.ListenAndServe(":8080", mux)
}
