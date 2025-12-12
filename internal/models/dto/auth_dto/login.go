package auth_dto

type LoginAuthRequest struct {
	Email 				string	`json:"email"`
	Password 			string	`json:"password"`
}