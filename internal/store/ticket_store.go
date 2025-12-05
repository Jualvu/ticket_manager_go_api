package store

import (
	"time"
	"log"
	"fmt"
	"github.com/jualvu/go-tickets-api/internal/models"
	"github.com/jualvu/go-tickets-api/internal/models/dto/ticket_dto"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// private
type ticketStore struct {
	db *sql.DB
}

// public
type TicketStore interface {
	GetAll() ([]*models.Ticket, error)
	GetByID(id int) (*models.Ticket, error)
	Create(u ticket_dto.CreateTicketRequest) (int64, error)
	Update(u ticket_dto.UpdateTicketRequest) error
	Delete(id ticket_dto.DeleteTicketRequest) error
}

// to create a new ticketStore struct
func NewTicketStore(db *sql.DB) (TicketStore) {
	us := ticketStore{db}
	return &us // -> return the ticket Store struct
}

func (s *ticketStore) GetAll() ([]*models.Ticket, error) {
	stmt, err := s.db.Prepare("SELECT id, title, description, state_id, priority_id, assigned_to_user_id, created_by_user_id, creation_date, last_update_date FROM tickets;")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer stmt.Close() // -> make sure to close stmt pointer

	var tickets []*models.Ticket

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		var curTicket models.Ticket
		err = rows.Scan(&curTicket.Id,
						&curTicket.Title,
						&curTicket.Description,
						&curTicket.State_id,
						&curTicket.Priority_id,
						&curTicket.Assigned_to_user_id,
						&curTicket.Created_by_user_id,
						&curTicket.Creation_date,
						&curTicket.Last_update_date)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		fmt.Printf("Found ticket: [%s], created at [%v]\n", curTicket.Title, curTicket.Creation_date)
		tickets = append(tickets, &curTicket)
	}
	return tickets, nil
}

func (s *ticketStore) GetByID(id int) (*models.Ticket, error) {
	stmt, err := s.db.Prepare("SELECT id, title, description, state_id, priority_id, assigned_to_user_id, created_by_user_id, creation_date, last_update_date FROM tickets AS t WHERE t.id = ?")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer stmt.Close()

	var ticket models.Ticket

	err = stmt.QueryRow(id).Scan(&ticket.Id,
								 &ticket.Title,
								 &ticket.Description,
								 &ticket.State_id,
								 &ticket.Priority_id,
								 &ticket.Assigned_to_user_id,
								 &ticket.Created_by_user_id,
								 &ticket.Creation_date,
								 &ticket.Last_update_date)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	fmt.Printf("Ticket with ID [%v] found: %v \n", id, ticket.Title)
	return &ticket, nil
}

func (s *ticketStore) Create(createTicketRequest ticket_dto.CreateTicketRequest) (int64, error) {
	stmt, err := s.db.Prepare("INSERT INTO tickets (title, description, state_id, priority_id, assigned_to_user_id, created_by_user_id, creation_date, last_update_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?);")
	if err != nil {
		log.Fatal(err)
		return -1, err
	}
	defer stmt.Close()

	creationDate := time.Now().Format("2006-01-02 15:04:05")

	res, err := stmt.Exec(createTicketRequest.Title,
						  createTicketRequest.Description,
						  createTicketRequest.State_id,
						  createTicketRequest.Priority_id,
						  createTicketRequest.Assigned_to_user_id,
						  createTicketRequest.Created_by_user_id,
						  creationDate,
						  creationDate)
	if err != nil {
		log.Fatal(err)
		return -1, err
	}

	newTicketID, err := res.LastInsertId() // return the id of the ticket we just added
	if err != nil {
		return -1, err
	}

	log.Printf("Succesfully created ticket with ID [%v]", newTicketID)
	return newTicketID, nil
}

func (s *ticketStore) Update(updateTicketRequest ticket_dto.UpdateTicketRequest) error {
	stmt, err := s.db.Prepare("UPDATE tickets AS t SET t.title = ?, t.description = ?, t.state_id = ?, t.priority_id = ?, t.assigned_to_user_id = ?, t.last_update_date = ? WHERE t.id = ?;")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()

	lastUpdateDate := time.Now().Format("2006-01-02 15:04:05")

	res, err := stmt.Exec(updateTicketRequest.Title,
						  updateTicketRequest.Description,
						  updateTicketRequest.State_id,
						  updateTicketRequest.Priority_id,
						  updateTicketRequest.Assigned_to_user_id,
						  lastUpdateDate,
						  updateTicketRequest.Id)
	if err != nil {
		log.Fatal(err)
		return err
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
		return err
	}

	if affectedRows == 0 {
		log.Printf("Warning: 0 rows updated giving ticket ID [%v]", updateTicketRequest.Id)
		return nil
	}

	log.Printf("Succesfully updated ticket with ID [%v]", updateTicketRequest.Id)

	return nil
}

func (s *ticketStore) Delete(deleteTicketRequest ticket_dto.DeleteTicketRequest) error {
	stmt, err := s.db.Prepare("DELETE FROM tickets AS t WHERE t.id = ?;")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(deleteTicketRequest.Id)
	if err != nil {
		log.Fatal(err)
		return err
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
		return err
	}
	
	if affectedRows == 0 {
		log.Printf("Warning: 0 rows deleted giving ticket with ID [%v]", deleteTicketRequest.Id)
		return nil
	}

	log.Printf("Succesfully deleted ticket with ID [%v]", deleteTicketRequest.Id)

	return nil
}


