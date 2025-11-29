package store

import (
	"time"
	"log"
	"fmt"
	"github.com/jualvu/go-tickets-api/internal/models"
	"github.com/jualvu/go-tickets-api/internal/models/dto/user_dto"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// private
type userStore struct {
	db *sql.DB
}

// public
type UserStore interface {
	GetAll() ([]*models.User, error)
	GetByID(id int) (*models.User, error)
	Create(u user_dto.CreateUserRequest) (int64, error)
	Update(u user_dto.UpdateUserRequest) error
	Delete(id user_dto.DeleteUserRequest) error
}

// to create a new userStore struct
func NewUserStore(db *sql.DB) (UserStore) {
	us := userStore{db} 
	return &us // -> return the user Store struct 
}

func (s *userStore) GetAll() ([]*models.User, error) {
	stmt, err := s.db.Prepare("SELECT id, name, email, rol_id, creation_date, last_update_date FROM users;") 
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer stmt.Close() // -> make sure to close stmt pointer	
	
	var users []*models.User

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		var curUser models.User
		err = rows.Scan(&curUser.Id,
						&curUser.Name,
						&curUser.Email,
						&curUser.Rol_id,
						&curUser.Creation_date,
						&curUser.Last_update_date)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		fmt.Printf("Found user: [%s], created at [%v]\n", curUser.Name, curUser.Creation_date)
		users = append(users, &curUser)
	}

	return users, nil
}

func (s *userStore) GetByID(id int) (*models.User, error) {
	stmt, err := s.db.Prepare("SELECT id, name, email, rol_id, creation_date, last_update_date FROM users AS u WHERE u.id = ?")
	if err != nil { 
		log.Fatal(err)	
		return nil, err
	}
	
	defer stmt.Close()
	
	var user models.User

	err = stmt.QueryRow(id).Scan(&user.Id,
								 &user.Name,
								 &user.Email,
								 &user.Rol_id,
								 &user.Creation_date,
								 &user.Last_update_date)

	if err != nil {
		log.Println(err)	
		return nil, err
	}

	fmt.Printf("User with ID [%v] found: %v \n", id, user.Name)
	return &user, nil
}

func (s *userStore) Create(createUserRequest user_dto.CreateUserRequest) (userID int64, err error) {
	stmt, err := s.db.Prepare("INSERT INTO users (name, email, password, rol_id, creation_date, last_update_date) VALUES (?, ?, ?, ?, ?, ?);")
	if err != nil {
		log.Fatal(err)	
		return -1, err
	}
	defer stmt.Close()

	creationDate := time.Now().Format("2006-01-02 15:04:05")

	res, err := stmt.Exec(createUserRequest.Name,
						  createUserRequest.Email,
						  createUserRequest.Password, 
						  createUserRequest.Rol_id,
						  creationDate,
						  creationDate)
	if err != nil {
		log.Fatal(err)	
		return -1, err
	}

	newUserID, err := res.LastInsertId() // return the id of the user we just added
	if err != nil {
		return -1, err
	}

	log.Printf("Succesfully added user:\nID [%v]\nName [%v]\n Email [%v]\n Password [%v]\n Rol_id [%v]\n", 
				newUserID,
				createUserRequest.Name, 
				createUserRequest.Email, 
				createUserRequest.Password, 
				createUserRequest.Rol_id)
	
	return newUserID, nil
}

func (s *userStore) Update(updateUserRequest user_dto.UpdateUserRequest) error {
	stmt, err := s.db.Prepare("UPDATE users SET name = ?, email = ?, password = ?, rol_id = ?, last_update_date = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)	
		return err
	}
	defer stmt.Close()

	lastUpdateDate := time.Now().Format("2006-01-02 15:04:05")

	res, err := stmt.Exec(updateUserRequest.Name, 
						  updateUserRequest.Email, 
						  updateUserRequest.Password, 
						  updateUserRequest.Rol_id, 
						  lastUpdateDate, 
						  updateUserRequest.Id)
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
		log.Print("Warning: 0 rows updated giving user ID [%v]", updateUserRequest.Id)
		return nil
	}

	log.Printf("Succesfully updated user with ID [%v]", updateUserRequest.Id)

	return nil
}

func (s *userStore) Delete(deleteUserRequest user_dto.DeleteUserRequest) error {
	stmt, err := s.db.Prepare("DELETE FROM users AS u WHERE u.id = ?")
	if err != nil {
		log.Fatal(err)	
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(deleteUserRequest.Id)
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
		log.Printf("Warning: 0 rows deleted giving user with ID [%v]", deleteUserRequest.Id)
		return nil
	}

	log.Printf("Succesfully deleted user with ID [%v]", deleteUserRequest.Id)

	return nil
}



