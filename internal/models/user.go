package models

type User struct {
	Id 					int		`json:"id"`	
	Name 				string	`json:"name"`	
	Email 				string	`json:"email"`
	Password 			string	`json:"password"`
	Rol_id 				int		`json:"rol_id"`
	Creation_date 		string	`json:"creation_date"`
	Last_update_date 	string	`json:"last_update_date"`
}