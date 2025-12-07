package store

import (
	"time"
	"log"
	"fmt"
	"github.com/jualvu/go-tickets-api/internal/models"
	"github.com/jualvu/go-tickets-api/internal/models/dto/comment_dto"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// private
type commentStore struct {
	db *sql.DB
}

// public
type CommentStore interface {
	GetAll() ([]*models.Comment, error)
	GetByID(id int) (*models.Comment, error)
	Create(u comment_dto.CreateCommentRequest) (int64, error)
	Update(u comment_dto.UpdateCommentRequest) error
	Delete(id comment_dto.DeleteCommentRequest) error
}

// to create a new commentStore struct
func NewCommentStore(db *sql.DB) (CommentStore) {
	us := commentStore{db} 
	return &us // -> return the comment Store struct 
}

func (s *commentStore) GetAll() ([]*models.Comment, error) {
	stmt, err := s.db.Prepare("SELECT id, content, created_by_user_id, ticket_id, creation_date, last_update_date FROM comments;") 
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer stmt.Close() // -> make sure to close stmt pointer	
	
	var comments []*models.Comment

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

		for rows.Next() {
		var curComment models.Comment
		err = rows.Scan(&curComment.Id,
						&curComment.Content,
						&curComment.Created_by_user_id,
						&curComment.Ticket_id,
						&curComment.Creation_date,
						&curComment.Last_update_date)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		fmt.Printf("Found comment: [%s], created at [%v]\n", curComment.Content, curComment.Creation_date)
		comments = append(comments, &curComment)
	}

	return comments, nil
}

func (s *commentStore) GetByID(id int) (*models.Comment, error) {
	stmt, err := s.db.Prepare("SELECT id, content, created_by_user_id, ticket_id, creation_date, last_update_date FROM comments AS c WHERE c.id = ?")
	if err != nil { 
		log.Fatal(err)	
		return nil, err
	}
	
	defer stmt.Close()
	
	var comment models.Comment

	err = stmt.QueryRow(id).Scan(&comment.Id,
								 &comment.Content,
								 &comment.Created_by_user_id,
								 &comment.Ticket_id,
								 &comment.Creation_date,
								 &comment.Last_update_date)

	if err != nil {
		log.Println(err)	
		return nil, err
	}

	fmt.Printf("Comment with ID [%v] found: %v \n", id, comment.Content)
	return &comment, nil
}

func (s *commentStore) Create(createCommentRequest comment_dto.CreateCommentRequest) (commentID int64, err error) {
	stmt, err := s.db.Prepare("INSERT INTO comments (content, created_by_user_id, ticket_id, creation_date, last_update_date) VALUES (?, ?, ?, ?, ?);")
	if err != nil {
		log.Fatal(err)	
		return -1, err
	}
	defer stmt.Close()

	creationDate := time.Now().Format("2006-01-02 15:04:05")

	res, err := stmt.Exec(createCommentRequest.Content,
						  createCommentRequest.Created_by_user_id,
						  createCommentRequest.Ticket_id,
						  creationDate,
						  creationDate)
	if err != nil {
		log.Fatal(err)	
		return -1, err
	}

	newCommentID, err := res.LastInsertId() // return the id of the comment we just added
	if err != nil {
		return -1, err
	}

	log.Printf("Succesfully added comment:\nID [%v]\nContent [%v]\nCreatedBy [%v]\nTicket [%v]\n",
				newCommentID,
				createCommentRequest.Content, 
				createCommentRequest.Created_by_user_id, 
				createCommentRequest.Ticket_id)

	return newCommentID, nil
}

func (s *commentStore) Update(updateCommentRequest comment_dto.UpdateCommentRequest) error {
	stmt, err := s.db.Prepare("UPDATE comments SET content = ?, last_update_date = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)	
		return err
	}
	defer stmt.Close()

	lastUpdateDate := time.Now().Format("2006-01-02 15:04:05")

	res, err := stmt.Exec(updateCommentRequest.Content, 
						  lastUpdateDate, 
						  updateCommentRequest.Id)
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
		log.Printf("Warning: 0 rows updated giving comment ID [%v]", updateCommentRequest.Id)
		return nil
	}

	log.Printf("Succesfully updated comment with ID [%v]", updateCommentRequest.Id)

	return nil
}


func (s *commentStore) Delete(deleteCommentRequest comment_dto.DeleteCommentRequest) error {
	stmt, err := s.db.Prepare("DELETE FROM comments AS c WHERE c.id = ?")
	if err != nil {
		log.Fatal(err)	
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(deleteCommentRequest.Id)
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
		log.Printf("Warning: 0 rows deleted giving comment with ID [%v]", deleteCommentRequest.Id)
		return nil
	}

	log.Printf("Succesfully deleted comment with ID [%v]", deleteCommentRequest.Id)

	return nil
}







