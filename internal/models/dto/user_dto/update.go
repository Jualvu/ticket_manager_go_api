package user_dto

type UpdateUserRequest struct {
	Id 					int		`json:"id"`	
	Name 				string	`json:"name"`	
	Email 				string	`json:"email"`
	Password 			string	`json:"password"`
	Rol_id 				int		`json:"rol_id"`
}